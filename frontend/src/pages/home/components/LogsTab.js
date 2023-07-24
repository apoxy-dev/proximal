import { useState, useEffect } from 'react';
import Box from '@mui/material/Box';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableFooter from '@mui/material/TableFooter';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import IconButton from '@mui/material/IconButton';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
import LinearProgress from '@mui/material/LinearProgress';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';

import Logs from '../../../api/Logs';

const TablePaginationActions = ({ count, page, rowsPerPage, onPageChange }) => {
  const handleFirstPageButtonClick = (event) => {
    onPageChange(event, 0);
  };

  const handleBackButtonClick = (event) => {
    onPageChange(event, page - 1);
  };

  const handleNextButtonClick = (event) => {
    onPageChange(event, page + 1);
  };

  const handleLastPageButtonClick = (event) => {
    onPageChange(event, Math.max(0, Math.ceil(count / rowsPerPage) - 1));
  };

  return (
    <Box sx={{ flexShrink: 0, ml: 2.5 }}>
      <IconButton
        onClick={handleFirstPageButtonClick}
        disabled={page === 0}
        aria-label="first page"
      >
        <FirstPageIcon />
      </IconButton>
      <IconButton
        onClick={handleBackButtonClick}
        disabled={page === 0}
        aria-label="previous page"
      >
        <KeyboardArrowLeft />
      </IconButton>
      <IconButton
        onClick={handleNextButtonClick}
        disabled={page >= Math.ceil(count / rowsPerPage) - 1}
        aria-label="next page"
      >
        <KeyboardArrowRight />
      </IconButton>
      <IconButton
        onClick={handleLastPageButtonClick}
        disabled={page >= Math.ceil(count / rowsPerPage) - 1}
        aria-label="last page"
      >
        <LastPageIcon />
      </IconButton>
    </Box>
  );
};

const GettingStarted = () => {
  return (
    <>
      <Container maxWidth="md">
        <Typography variant="h4" align="center" marginTop={2}>No Logs</Typography>
        <Typography variant="body1" align="center" marginTop={2}>
          Place a request to the Envoy server listening on <strong>port 18000</strong> and then logs will appear here.
        </Typography>
        <code style={{
          display: 'block',
          margin: '1rem',
          textAlign: 'center',
        }}>
          curl http://localhost:18000
        </code>
      </Container>
    </>
  );
};

const LogsTab = () => {
  const [ page, setPage ] = useState(0);
  const [ rowsPerPage, setRowsPerPage ] = useState(10);
  const [ loading, setLoading ] = useState(true);
  const [ logs, setLogs ] = useState([]);

  let reloadTimeout = null;

  const load = async () => {
    let r = await Logs.List();
    if (r.logs?.length > 0) {
      setLogs(r.logs);
      // Complete the tutorial since they've sent a request.
      localStorage.setItem('proximalTutorialComplete', true);
    }
    setLoading(false);
    reloadTimeout = setTimeout(load, 1000);
  };
  useEffect(() => {
    load();
    return () => clearTimeout(reloadTimeout);
  }, []);

  if (loading) {
    return <LinearProgress color="secondary" sx={{ marginTop: '2rem' }} />;
  }

  if (logs.length === 0) {
    return <GettingStarted />;
  }

  // Avoid a layout jump when reaching the last page with empty rows.
  const emptyRows =
    page > 0 ? Math.max(0, (1 + page) * rowsPerPage - logs.length) : 0;

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  return (
    <TableContainer component={Box}>
      <Table sx={{ minWidth: 500 }} aria-label="custom pagination table">
        <TableHead>
          <TableRow>
            <TableCell>Timestamp</TableCell>
            <TableCell align="right">Method</TableCell>
            <TableCell align="left">Path</TableCell>
            <TableCell align="right">Response Code</TableCell>
            <TableCell align="right">Duration</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {(logs > 0
            ? logs.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
            : logs
          ).map((row) => (
            <TableRow key={row.id}>
              <TableCell component="th" scope="row">
                {row.timestamp}
              </TableCell>
              <TableCell style={{ width: 160 }} align="right">
                {row.http.request.requestMethod}
              </TableCell>
              <TableCell align="left">
                {row.http.request.path}
              </TableCell>
              <TableCell style={{ width: 160 }} align="right">
                {row.http.response.responseCode}
              </TableCell>
              <TableCell style={{ width: 160 }} align="right">
                {row.http.commonProperties.duration}
              </TableCell>
            </TableRow>
          ))}
          {emptyRows > 0 && (
            <TableRow style={{ height: 53 * emptyRows }}>
              <TableCell colSpan={7} />
            </TableRow>
          )}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TablePagination
              rowsPerPageOptions={[5, 10, 25, { label: 'All', value: -1 }]}
              colSpan={5}
              count={logs.length}
              rowsPerPage={rowsPerPage}
              page={page}
              SelectProps={{
                inputProps: {
                  'aria-label': 'rows per page',
                },
                native: true,
              }}
              onPageChange={handleChangePage}
              onRowsPerPageChange={handleChangeRowsPerPage}
              ActionsComponent={TablePaginationActions}
            />
          </TableRow>
        </TableFooter>
      </Table>
    </TableContainer>
  );
};

export default LogsTab;
