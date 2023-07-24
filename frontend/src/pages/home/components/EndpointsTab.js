import { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import Button from '@mui/material/Button';
import Checkbox from '@mui/material/Checkbox';
import Grid from '@mui/material/Grid';
import FormControlLabel from '@mui/material/FormControlLabel';
import LinearProgress from '@mui/material/LinearProgress';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import StatusIndicator from './StatusIndicator';
import Modal from '@mui/material/Modal';
import MenuList from '@mui/material/MenuList';
import MenuItem from '@mui/material/MenuItem';
import SvgIcon from '@mui/material/SvgIcon';
import {
  ArrowRight,
  CloudAdd,
  CloudCross,
  CloudConnection,
} from 'iconsax-react';

import SimpleModal from '../../../components/SimpleModal';
import TextInput from '../../../components/TextInput';
import InputExample from '../../../components/InputExample';
import Endpoints from '../../../api/Endpoints';

const EndpointsStatus = (props) => {
  if (props.status === 'waiting') {
    return (
      <StatusIndicator icon={<CloudCross />} variant="disabled">
        No Endpoints Configured
      </StatusIndicator>
    );
  } else if (props.status === 'ready') {
    return (
      <StatusIndicator icon={<CloudAdd />}>
        Endpoints Ready
      </StatusIndicator>
    );
  }
  return <></>;
};

const CustomEndpoint = (props) => {
  const [request, setRequest] = useState({});

  const onChange = (key) => {
    return (e) => {
      setRequest({ ...request, [key]: e.target.value });
    };
  };
  const onTLSToggle = (e) => {
    setRequest({ ...request, 'use_tls': e.target.checked });
  };

  const create = async () => {
    let endpoint = {
      cluster: request['cluster'],
      addresses: [
        {
          host: request['host'],
          port: Number(request['port']),
        },
      ],
      use_tls: request['use_tls'],
    };
    props.onClose();
    await Endpoints.Create(endpoint);
    props.reload(endpoint.cluster);
  };

  const valid = () => {
    if (!request['cluster'] || request['cluster'].length < 1 || request['cluster'].includes(' ')) {
      return false;
    }
    if (!request['host'] || request['host'].length < 2) {
      return false;
    }
    if (!request['port'] || ''+Number(request['port']) !== request['port']) {
      return false;
    }
    return true;
  };

  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Add an Endpoint</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center">
            <TextInput placeholder="Cluster Name" sx={{ width: '100%' }} onChange={onChange('cluster')} />
            <InputExample>eg. <strong>local_dev</strong> or <strong>httpbin</strong></InputExample>
            <TextInput placeholder="Host" sx={{ width: '100%' }} onChange={onChange('host')} />
            <InputExample>eg. <strong>127.0.0.1</strong> or <strong>httpbin.org</strong></InputExample>
            <TextInput placeholder="Port" sx={{ width: '100%' }} onChange={onChange('port')} />
            <InputExample>eg. <strong>80</strong> or <strong>443</strong></InputExample>
          </Grid>
          <Grid item xs={12} container alignItems="center" marginTop={-4}>
            <FormControlLabel control={<Checkbox checked={request['use_tls']} onChange={onTLSToggle} />} label="Use TLS" />
          </Grid>
          <Grid item xs={12} container>
            <Button
              variant="contained"
              endIcon={<ArrowRight />}
              sx={{ margin: '0 auto' }}
              disabled={!valid()}
              onClick={create}
              data-ph-capture-attribute-added-address={`${request['host']}:${request['port']}`}
            >
              Add It
            </Button>
          </Grid>
        </Grid>
      </SimpleModal>
    </Modal>
  );
};

const GettingStarted = (props) => {
  return (
    <>
      <Container maxWidth="md">
        <Typography variant="h4" align="center" marginTop={2}>No Endpoints Configured</Typography>
        <Typography variant="body1" align="center" marginTop={2}>
          Endpoints are hosts that handle requests from Envoy and return responses.
        </Typography>
      </Container>
      <Grid container spacing={4} marginTop={2}>
        <Grid item xs={4} align="right">
          <Button variant="contained" onClick={() => props.useExample('httpbin')}>
            Use httpbin.org
          </Button>
        </Grid>
        <Grid item xs={8} container alignItems="center">
          <Typography variant="body1">
            <a href="http://httpbin.org" target="_blank">httpbin.org</a> is a simple HTTP server with a variety of test endpoints.
          </Typography>
        </Grid>

        <Grid item xs={4} align="right">
          <Button variant="contained" onClick={() => props.useExample('httpstatus')}>
            Use httpstat.us
          </Button>
        </Grid>
        <Grid item xs={8} container alignItems="center">
          <Typography variant="body1">
            <a href="http://httpstat.us" target="_blank">httpstat.us</a> is a super simple service for generating different HTTP codes.
          </Typography>
        </Grid>

        <Grid item xs={4} align="right">
          <Button variant="contained" onClick={props.openPopup}>
            Add your own
          </Button>
        </Grid>
        <Grid item xs={8} container alignItems="center">
          <Typography variant="body1">
            Use your own server as an upstream.
          </Typography>
        </Grid>

      </Grid>
    </>
  );
};

const EndpointSettings = ({ onlyEndpoint, endpoint, reload }) => {
  if (!endpoint) {
    return <>Loading...</>;
  }
  const Detail = ({ name, description }) => (
    <>
      <Grid item xs={4}>
        <code>
          {name}
        </code>
      </Grid>
      <Grid item xs={8}>
        <code>
          {description}
        </code>
      </Grid>
    </>
  );
  const deleteEndpoint = async () => {
    await Endpoints.Delete(endpoint.cluster);
    reload('');
  };
  const defaultEndpoint = async () => {
    await Endpoints.SetDefault(endpoint);
    reload('');
  };
  return <>
    <Grid container marginTop={0} spacing={2}>
      <Grid item xs={12} alignItems="left" align="left" marginBottom={2}>
        <Button
          color="primary"
          variant="outlined"
          onClick={deleteEndpoint}
          disabled={onlyEndpoint || endpoint.defaultUpstream}
        >
          Delete Endpoint
        </Button>
        <Button
          color="primary"
          variant="outlined"
          sx={{ marginLeft: '16px' }}
          disabled={endpoint.defaultUpstream}
          onClick={defaultEndpoint}
        >
          Make Default
        </Button>
      </Grid>
      <Detail name="Cluster Name" description={endpoint.cluster} />
      <Detail name="Default Upstream?" description={endpoint.defaultUpstream ? 'yes' : 'no'} />
      <Detail name="Use TLS?" description={endpoint.useTls ? 'yes' : 'no'} />
      <Detail name="Addresses" description={endpoint.addresses.map((a) => {
        return `${a.host}:${a.port}\n`;
      })} />
    </Grid>
  </>;
};

const FirstEndpointTutorial = (props) => {
  const navigate = useNavigate();
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Next Up: Logs</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center" justifyContent="center">
            <Typography variant="body1">
              Awesome. We're ready to send Envoy some traffic.
            </Typography>
          </Grid>
          <Grid item xs={12} container>
            <Button variant="contained" endIcon={<ArrowRight />} sx={{ margin: '0 auto' }} onClick={() => navigate('/logs')}>
              Continue
            </Button>
          </Grid>
        </Grid>
      </SimpleModal>
    </Modal>
  );
};

const EndpointStatus = () => {
  return <SvgIcon sx={{
    marginRight: '16px',
    width: '24px',
    height: '24px',
  }}>
    <CloudConnection />
  </SvgIcon>;
};

const EndpointsTab = () => {
  const [ endpoints, setEndpoints ] = useState([]);
  const [ active, setActive ] = useState('');
  const [ loading, setLoading ] = useState(true);
  const [ open, setOpen ] = useState(false);
  const location = useLocation();
  const navigate = useNavigate();
  const tutorialActive = location.search.includes('tutorial');

  const load = async () => {
    let r = await Endpoints.List();
    if (r.endpoints?.length > 0) {
      setEndpoints(r.endpoints);
      if (active === '') {
        setActive(r.endpoints[0].cluster);
      }
    } else {
      setEndpoints([]);
    }
    setLoading(false);
  };
  useEffect(() => {
    load();
  }, []);

  const endpointByCluster = (cluster) => {
    return endpoints.find(e => e.cluster === cluster);
  };

  const reload = async (cluster) => {
    setLoading(true);
    if (cluster) {
      setActive(cluster);
    }
    await load();
  };

  const useExample = async (svc) => {
    if (svc === 'httpbin') {
      await Endpoints.Create({
        cluster: 'httpbin',
        addresses: [
          {
            host: 'httpbin.org',
            port: 80,
          },
        ],
      });
    } else if (svc === 'httpstatus') {
      await Endpoints.Create({
        cluster: 'httpstatus',
        addresses: [
          {
            host: 'httpstat.us',
            port: 80,
          },
        ],
      });
    }
    reload();
    // Show the tutorial popup if they haven't completed it yet.
    const tutorialComplete = !!localStorage.getItem('proximalTutorialComplete');
    if (!tutorialComplete) {
      navigate('/endpoints?tutorial=1');
    }
  };

  if (loading) {
    return <LinearProgress color="secondary" sx={{ marginTop: '2rem' }} />;
  }

  if (endpoints.length === 0) {
    return (
      <>
        <CustomEndpoint open={open} onClose={() => setOpen(false) } reload={reload} />
        <GettingStarted openPopup={() => setOpen(true)} useExample={useExample} />
      </>
    );
  }
  return (
    <>
      <CustomEndpoint open={open} onClose={() => setOpen(false) } reload={reload} />
      <FirstEndpointTutorial open={tutorialActive} onClose={() => navigate('/endpoints') } />
      <Grid container sx={{ paddingBottom: 2, borderBottom: 1, borderColor: 'divider' }}>
        <Grid item container xs={6} alignItems="center">
          <EndpointsStatus status={'ready'} />
        </Grid>
        <Grid item xs={6} alignItems="right" align="right">
          <Button color="secondary" variant="contained" onClick={() => setOpen(true)}>
            Add an Endpoint
          </Button>
        </Grid>
      </Grid>
      <Grid container spacing={2}>
        <Grid item xs={3} marginTop={1}>
          <MenuList sx={{ width: '100%' }}>
            {endpoints.map((m) => (
              <MenuItem
                key={m.cluster}
                selected={m.cluster === active}
                onClick={() => setActive(m.cluster)}
              >
                <EndpointStatus status={m.status} />
                {m.cluster}
              </MenuItem>
            ))}
          </MenuList>
        </Grid>
        <Grid item xs={9}>
          <EndpointSettings
            onlyEndpoint={endpoints.length === 1}
            endpoint={endpointByCluster(active)}
            reload={reload} />
        </Grid>
      </Grid>
    </>
  );
};

export default EndpointsTab;
