import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import Box from '@mui/material/Box';

import { useLocation, useNavigate } from 'react-router-dom';

import CustomTabPanel from './components/CustomTabPanel';
import ExtensionsTab from './components/ExtensionsTab';
import EndpointsTab from './components/EndpointsTab';
import LogsTab from './components/LogsTab';

const a11yProps = (index) => {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
};

const Home = ()	=> {
  // TODO: const [extState, setExtState] = useState('waiting');
  // TODO: const [endpointsState, setEndpointsState] = useState('waiting');

  const tabs = {
    0: '/extensions',
    1: '/endpoints',
    2: '/logs',
  };
  let paths = {};
  Object.keys(tabs).forEach((key) => {
    paths[tabs[key]] = Number(key);
  });

  const location = useLocation();
  const value = paths[location.pathname];

  const navigate = useNavigate();
  const handleChange = (event, newValue) => {
    navigate(tabs[newValue]);
  };

  return (
    <Box sx={{ width: '100%' }}>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={value} onChange={handleChange} aria-label="basic tabs example" indicatorColor="secondary">
          <Tab label="Extensions" {...a11yProps(0)} />
          <Tab label="Endpoints" {...a11yProps(1)} />
          <Tab label="Logs" {...a11yProps(2)} />
        </Tabs>
      </Box>
      <CustomTabPanel value={value} index={0}>
        <ExtensionsTab />
      </CustomTabPanel>
      <CustomTabPanel value={value} index={1}>
        <EndpointsTab />
      </CustomTabPanel>
      <CustomTabPanel value={value} index={2}>
        <LogsTab />
      </CustomTabPanel>
    </Box>
  );
};

export default Home;
