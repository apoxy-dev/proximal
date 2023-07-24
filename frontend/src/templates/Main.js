import * as React from 'react';
import { Outlet } from 'react-router-dom';
import PropTypes from 'prop-types';
import { useTheme } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import CssBaseline from '@mui/material/CssBaseline';
import useScrollTrigger from '@mui/material/useScrollTrigger';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import GithubIcon from '@mui/icons-material/GitHub';
import { ArrangeHorizontalCircle } from 'iconsax-react';
import { Link } from 'react-router-dom';

function ElevationScroll(props) {
  const { children, window } = props;
  // Note that you normally won't need to set the window ref as useScrollTrigger
  // will default to window.
  // This is only being set here because the demo is in an iframe.
  const trigger = useScrollTrigger({
    disableHysteresis: true,
    threshold: 0,
    target: window ? window() : undefined,
  });

  return React.cloneElement(children, {
    elevation: trigger ? 4 : 0,
  });
}

ElevationScroll.propTypes = {
  children: PropTypes.element.isRequired,
  /**
   * Injected by the documentation to work in an iframe.
   * You won't need it on your project.
   */
  window: PropTypes.func,
};

const Main = (props) => {
  const theme = useTheme();
  return (
    <>
      <CssBaseline />
      <ElevationScroll {...props}>
        <AppBar>
          <Toolbar>
            <Box sx={{ display: 'flex', flexGrow: 1, alignItems: 'center' }}>
              <Box sx={{ display: 'flex', padding: '0 8px 0 0' }}>
                <Link to="/" style={{ display: 'flex' }}>
                  <ArrangeHorizontalCircle
                    size="35"
                    color={theme.palette.secondary.main}
                  />
                </Link>
              </Box>
              <Typography variant="h5" component="div">
              Proximal
              </Typography>
            </Box>
            <Link
              to="https://github.com/apoxy-dev/proximal"
              target="_blank"
            >
              <Button
                variant="outlined"
                color="secondary"
                endIcon={<GithubIcon />}
              >
                View on GitHub
              </Button>
            </Link>
          </Toolbar>
        </AppBar>
      </ElevationScroll>
      <Toolbar />
      <Container sx={{ marginTop: 2 }}>
        <Outlet />
      </Container>
    </>
  );
};

export default Main;
