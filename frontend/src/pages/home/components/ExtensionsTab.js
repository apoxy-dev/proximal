import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import MenuList from '@mui/material/MenuList';
import MenuItem from '@mui/material/MenuItem';
import StatusIndicator from './StatusIndicator';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import Container from '@mui/material/Container';
import LinearProgress from '@mui/material/LinearProgress';
import Modal from '@mui/material/Modal';
import Link from '@mui/material/Link';
import Select from '@mui/material/Select';
import InputLabel from '@mui/material/InputLabel';
import FormControl from '@mui/material/FormControl';
import {
  ArrowRight,
  FlashCircle,
  Refresh,
  Forbidden,
} from 'iconsax-react';
import SvgIcon from '@mui/material/SvgIcon';

import Endpoints from '../../../api/Endpoints';
import Middleware from '../../../api/Middleware';
import TextInput from '../../../components/TextInput';
import SimpleModal from '../../../components/SimpleModal';
import InputExample from '../../../components/InputExample';

const Examples = [
  {
    slug: 'http_headers',
    language: 'Rust',
    link: 'https://github.com/proxy-wasm/proxy-wasm-rust-sdk/tree/master/examples/http_headers',
    description: 'Logs all of the incoming and outgoing HTTP headers.',
    request: {
      slug: 'http_headers',
      ingest_params: {
        type: 'GITHUB',
        github_repo: 'https://github.com/proxy-wasm/proxy-wasm-rust-sdk.git',
        branch: 'master',
        language: 'RUST',
        build_args: [
          '--manifest-path',
          './examples/http_headers/Cargo.toml',
        ],
      },
    },
  },
  {
    slug: 'http_auth_random',
    language: 'Go',
    link: 'https://github.com/tetratelabs/proxy-wasm-go-sdk/tree/main/examples/http_auth_random',
    description: 'Authorizes or denies requests based on a random number returned from a remote HTTP server.',
    request: {
      slug: 'http_auth_random',
      ingest_params: {
        type: 'GITHUB',
        github_repo: 'https://github.com/tetratelabs/proxy-wasm-go-sdk.git',
        branch: 'main',
        language: 'GO',
        build_args: [
          './examples/http_auth_random/main.go',
        ],
      },
    },
    upstreams: [
      {
        cluster: 'httpbin',
        host: 'httpbin.org',
        port: 80,
      },
    ],
  },
];

const Example = ({ example, onEndpoints }) => {
  const tryExample = async () => {
    const middlewares = await Middleware.Create(example.request);
    onEndpoints(middlewares, example.upstreams);
  };
  return (
    <>
      <Grid item xs={8} key={example.slug}>
        <Link href={example.link} target="_blank">
          <Typography variant="h6" align="right">
            {example.slug} - {example.language}
          </Typography>
        </Link>
        <Typography variant="body1" align="right">
          {example.description}
        </Typography>
      </Grid>
      <Grid item xs={4} container alignItems="center" justifyContent="center">
        <Button variant="contained" endIcon={<ArrowRight />} onClick={tryExample}>
          Try It
        </Button>
      </Grid>
      <Grid item xs={12}>
        <Divider />
      </Grid>
    </>
  );
};

const ExtensionsStatus = (props) => {
  if (props.status === 'waiting') {
    return (
      <StatusIndicator icon={<FlashCircle />} variant="disabled">
        Awaiting Configuration...
      </StatusIndicator>
    );
  } else if (props.status === 'ready') {
    return (
      <StatusIndicator icon={<FlashCircle />}>
        Watching for changes...
      </StatusIndicator>
    );
  } else if (props.status === 'building') {
    return (
      <StatusIndicator icon={<FlashCircle />}>
        Building...
      </StatusIndicator>
    );
  }
  return <></>;
};

const AddExtensions = (props) => (
  <Grid container spacing={4} marginTop={2}>
    <Grid item xs={4} align="right">
      <Button variant="contained" onClick={props.openPopup('example')}>
        Try an Example
      </Button>
    </Grid>
    <Grid item xs={8} container alignItems="center">
      <Typography variant="body1">
        Browse example extensions we've curated from the community.
      </Typography>
    </Grid>

    <Grid item xs={4} align="right">
      <Button variant="contained" onClick={props.openPopup('github')}>
        From GitHub
      </Button>
    </Grid>
    <Grid item xs={8} container alignItems="center">
      <Typography variant="body1">
        Compile from a GitHub repository containing code for an extension.
      </Typography>
    </Grid>

    <Grid item xs={4} align="right">
      <Button variant="contained" onClick={props.openPopup('local')}>
        Local Repo
      </Button>
    </Grid>
    <Grid item xs={8} container alignItems="center">
      <Typography variant="body1">
        Compile code in a local directory. Proximal will automatically watch for changes.
      </Typography>
    </Grid>
  </Grid>
);

const GettingStarted = (props) => {
  return (
    <>
      <Container maxWidth="md">
        <Typography variant="h4" align="center" marginTop={2}>No Extensions Configured</Typography>
        <Typography variant="body1" align="center" marginTop={2}>
          Extensions are <a href="https://github.com/proxy-wasm/spec" target="_blank">proxy-wasm based</a> WebAssembly modules that extend the
          functionality of your application.
        </Typography>
      </Container>
      <AddExtensions openPopup={props.openPopup} />
    </>
  );
};

const FromExample = (props) => {
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Example Extensions</Typography>
        <Grid container spacing={2} marginTop={1}>
          <Grid item xs={12}>
            <Divider />
          </Grid>
          {Examples.map((example) => (
            <Example key={example.slug} example={example} onEndpoints={props.onEndpoints} />
          ))}
        </Grid>
        <Typography variant="body1" align="center" marginTop={4}>
          Want us to add more examples or languages?<br/>
          Let us know by emailing <Link href="mailto:hello@apoxy.dev">hello@apoxy.dev</Link>.
        </Typography>
      </SimpleModal>
    </Modal>
  );
};

const FromGitHub = (props) => {
  const [language, setLanguage] = useState('');
  const [request, setRequest] = useState({});

  const onChange = (key) => {
    return (e) => {
      let value = e.target.value;
      if (key === 'ref') {
        const commitRegex = /^[0-9a-f]{40}$/i;
        if (value.match(commitRegex)) {
          key = 'commit';
        } else {
          key = 'branch';
        }
      }
      if (key === 'args') {
        value = value === '' ? [] : value.split(' ');
      }
      setRequest({ ...request, [key]: value });
    };
  };

  const create = async () => {
    let middleware = {
      slug: request['slug'],
      ingest_params: {
        type: 'GITHUB',
        github_repo: request['repo'],
        branch: request['branch'],
        commit: request['commit'],
        language: language,
      }
    };
    if (request['args']) {
      middleware.ingest_params['build_args'] = request['args'];
    }
    props.onClose();
    props.onEndpoints(await Middleware.Create(middleware), []);
  };

  const valid = () => {
    if (language === '') {
      return false;
    }
    if (!request['slug'] || request['slug'].length < 2 || request['slug'].includes(' ')) {
      return false;
    }
    if (!request['repo'] || request['repo'].length < 10) {
      return false;
    }
    if (!request['branch'] && !request['commit']) {
      return false;
    }
    return true;
  };
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">GitHub Extension</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center">
            <TextInput placeholder="Slug" sx={{ width: '100%' }} onChange={onChange('slug')} />
            <InputExample>eg. <strong>my_extension</strong></InputExample>
            <TextInput placeholder="Git Repo" sx={{ width: '100%' }} onChange={onChange('repo')} />
            <InputExample>eg. <strong>https://github.com/proxy-wasm/proxy-wasm-rust-sdk.git</strong></InputExample>
            <TextInput placeholder="Branch / Commit" sx={{ width: '100%' }} onChange={onChange('ref')} />
            <InputExample>eg. <strong>main</strong> or <strong>&lt;commit sha&gt;</strong></InputExample>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">Language</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                label="Language"
                sx={{ width: '100%' }}
                value={language}
                onChange={ (e) => setLanguage(e.target.value) }
              >
                <MenuItem value="GO">Go</MenuItem>
                <MenuItem value="RUST">Rust</MenuItem>
              </Select>
            </FormControl>
            <TextInput placeholder="Build Args (optional)" sx={{ width: '100%' }} onChange={onChange('args')} />
            <InputExample>flags to pass to the compiler eg. <strong>--manifest-path ./subdir/Cargo.toml</strong></InputExample>
          </Grid>
          <Grid item xs={12} container>
            <Button variant="contained" endIcon={<ArrowRight />} sx={{ margin: '0 auto' }} onClick={create} disabled={!valid()}>
              Let's Go
            </Button>
          </Grid>
        </Grid>
        <Typography variant="body1" align="center" marginTop={4}>
          Think we should add this to our examples?<br/>
          Let us know by emailing <Link href="mailto:hello@apoxy.dev">hello@apoxy.dev</Link>.
        </Typography>
      </SimpleModal>
    </Modal>
  );
};

const FromLocalRepo = (props) => {
  const [language, setLanguage] = useState('');
  const [request, setRequest] = useState({});

  const onChange = (key) => {
    return (e) => {
      let value = e.target.value;
      if (key === 'args') {
        value = value === '' ? [] : value.split(' ');
      }
      setRequest({ ...request, [key]: value });
    };
  };

  const create = async () => {
    let middleware = {
      slug: request['slug'],
      ingest_params: {
        type: 'DIRECT',
        language: language,
        watch_dir: request['watch_dir'],
      }
    };
    if (request['args']) {
      middleware.ingest_params['build_args'] = request['args'];
    }
    props.onClose();
    props.onEndpoints(await Middleware.Create(middleware), []);
  };

  const valid = () => {
    if (language === '') {
      return false;
    }
    if (!request['slug'] || request['slug'].length < 2 || request['slug'].includes(' ')) {
      return false;
    }
    if (!request['watch_dir'] || request['watch_dir'].length < 2) {
      return false;
    }
    return true;
  };

  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Local Development</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center">
            <TextInput placeholder="Slug" sx={{ width: '100%' }} onChange={onChange('slug')} />
            <InputExample>eg. <strong>my_extension</strong></InputExample>
            <TextInput placeholder="Host Path" sx={{ width: '100%' }} onChange={onChange('watch_dir')} />
            <InputExample>
              <strong>Absolute path</strong> on the server's host<br/>
              eg. <strong>/home/bobross/go/src/github.com/happy-accidents/proxy-art</strong>
            </InputExample>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">Language</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                label="Language"
                sx={{ width: '100%' }}
                value={language}
                onChange={ (e) => setLanguage(e.target.value) }
              >
                <MenuItem value="GO">Go</MenuItem>
                <MenuItem value="RUST">Rust</MenuItem>
              </Select>
            </FormControl>
            <TextInput placeholder="Build Args (optional)" sx={{ width: '100%' }} onChange={onChange('args')} />
            <InputExample>flags to pass to the compiler eg. <strong>--manifest-path ./subdir/Cargo.toml</strong></InputExample>
          </Grid>
          <Grid item xs={12} container>
            <Button variant="contained" endIcon={<ArrowRight />} sx={{ margin: '0 auto' }} onClick={create} disabled={!valid()}>
              Let's Go
            </Button>
          </Grid>
        </Grid>
      </SimpleModal>
    </Modal>
  );
};

const AddModal = (props) => {
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Add an Extension</Typography>
        <AddExtensions openPopup={props.openPopup} />
      </SimpleModal>
    </Modal>
  );
};

const FirstExtensionTutorial = (props) => {
  const navigate = useNavigate();
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Next Up: Endpoints</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center" justifyContent="center">
            <Typography variant="body1">
              Awesome. We've added an extension now we need to configure an endpoint.
            </Typography>
          </Grid>
          <Grid item xs={12} container>
            <Button variant="contained" endIcon={<ArrowRight />} sx={{ margin: '0 auto' }} onClick={() => navigate('/endpoints')}>
              Continue
            </Button>
          </Grid>
        </Grid>
      </SimpleModal>
    </Modal>
  );
};

const AddEndpoints = (props) => {
  const navigate = useNavigate();
  const addAndContinue = async () => {
    for (let i = 0; i < props.endpoints.length; i++) {
      await Endpoints.Create({
        cluster: props.endpoints[i].cluster,
        addresses: [
          {
            host: props.endpoints[i].host,
            port: props.endpoints[i].port,
          },
        ],
      });
    }
    const tutorialComplete = !!localStorage.getItem('proximalTutorialComplete');
    if (!tutorialComplete) {
      navigate('/endpoints?tutorial=1');
    } else{
      navigate('/endpoints');
    }
  };
  const Endpoint = ({ endpoint }) => (
    <>
      <strong>{endpoint.cluster}</strong>
      <ArrowRight style={{
        margin: '0 8px',
        position: 'relative',
        top: '6px',
      }}/>
      {endpoint.host}:{endpoint.port}
      <br/>
    </>
  );
  if (!props.endpoints || props.endpoints.length === 0) {
    return <FirstExtensionTutorial open={props.open} onClose={props.onClose} />;
  }
  return (
    <Modal open={props.open} onClose={props.onClose}>
      <SimpleModal>
        <Typography variant="h4" align="center">Add Required Endpoints</Typography>
        <Grid container spacing={4} marginTop={1}>
          <Grid item xs={12} container alignItems="center" justifyContent="center">
            <Typography variant="body1" align="center">
              This example has required endpoints: <br/>
              {props.endpoints.map((endpoint) => (
                <Endpoint key={endpoint.cluster} endpoint={endpoint} />
              ))}
            </Typography>
          </Grid>
          <Grid item xs={12} container>
            <Button variant="contained" endIcon={<ArrowRight />} sx={{ margin: '0 auto' }} onClick={addAndContinue}>
              Add and continue
            </Button>
          </Grid>
        </Grid>
      </SimpleModal>
    </Modal>
  );
};

const ExtensionStatus = (props) => {
  let icon = <FlashCircle />;
  let sx = {};
  if (props.status === 'PENDING') {
    icon = <Refresh />;
    sx = {
      animation: 'spin 2s linear infinite',
      '@keyframes spin': {
        '0%': {
          transform: 'rotate(0deg)',
        },
        '100%': {
          transform: 'rotate(360deg)',
        },
      },
    };
  }
  if (props.status === 'ERROR') {
    icon = <Forbidden />;
    sx = {
      color: 'red',
    };
  }
  return <SvgIcon sx={{
    marginRight: '16px',
    width: '24px',
    height: '24px',
    ...sx,
  }}>
    {icon}
  </SvgIcon>;
};

const ExtensionSettings = ({ extension, reload }) => {
  let details = {
    'Status': extension.status,
    'Slug': extension.slug,
    'Language': extension.ingestParams.language,
    'Type': extension.ingestParams.type,
  };
  let rebuildButton = <></>;

  const rebuild = async () => {
    await Middleware.TriggerBuild(extension.slug);
    reload();
  };
  switch (extension.ingestParams.type) {
    case 'GITHUB':
      details['Repo'] = <Link href={extension.ingestParams.githubRepo} target="_blank" rel="noreferrer">{extension.ingestParams.githubRepo}</Link>;
      if (details['Branch']) {
        details['Branch'] = extension.ingestParams.branch;
      } else {
        details['Commit'] = extension.ingestParams.commit;
      }
      break;
    case 'DIRECT':
      details['Watch Dir'] = extension.ingestParams.watchDir;
      rebuildButton = (
        <Button color="primary" variant="outlined" onClick={rebuild} sx={{ marginLeft: '16px' }}>
          Rebuild
        </Button>
      );
      break;
  }
  details['Build Args'] = JSON.stringify(extension.ingestParams.buildArgs);
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
  const deleteExtension = async () => {
    await Middleware.Delete(extension.slug);
    reload();
  };
  return <>
    <Grid container marginTop={0} spacing={2}>
      <Grid item xs={12} alignItems="left" align="left" marginBottom={2}>
        <Button color="primary" variant="outlined" onClick={deleteExtension}>
          Delete Extension
        </Button>
        {rebuildButton}
      </Grid>
      {Object.keys(details).map((k, index) => (
        <Detail
          name={k}
          description={details[k]}
          key={index}
        />
      ))
      }
    </Grid>
  </>;
};

const ExtensionsTab = () => {
  let [ loading, setLoading ] = useState(true);
  let [ middlewares, setMiddlewares ] = useState([]);
  let [ popup, setPopup ] = useState('');
  let [ active, setActive ] = useState('');
  let [ endpoints, setEndpoints ] = useState([]);

  const closePopup = () => {
    setPopup('');
  };

  const openPopup = (type) => {
    return () => {
      setPopup(type);
    };
  };

  const onEndpoints = (middleware, endpoints) => {
    middlewares.push(middleware);
    setMiddlewares(middlewares);
    setEndpoints(endpoints);
    setActive(middleware.slug);
    const tutorialComplete = !!localStorage.getItem('proximalTutorialComplete');
    if (!tutorialComplete || endpoints?.length > 0) {
      openPopup('endpoints')();
    } else {
      closePopup();
    }
  };

  let reloadTimeout = null;

  const load = async () => {
    let r = await Middleware.List();
    if (r.middlewares?.length > 0) {
      setMiddlewares(r.middlewares);
      if (active === '') {
        setActive(r.middlewares[0].slug);
      }
    } else {
      setMiddlewares([]);
    }
    setLoading(false);
    reloadTimeout = setTimeout(load, 1000);
  };
  useEffect(() => {
    load();
    return () => {
      clearTimeout(reloadTimeout);
    };
  }, []);

  const reload = async () => {
    setLoading(true);
    await load();
  };

  const middlewareBySlug = (slug) => {
    return middlewares.find(m => m.slug === slug);
  };

  if (loading) {
    return <LinearProgress color="secondary" sx={{ marginTop: '2rem' }} />;
  }

  if (middlewares.length === 0) {
    return (
      <>
        <FromExample open={popup === 'example'} onClose={closePopup} onEndpoints={onEndpoints} />
        <FromGitHub open={popup === 'github'} onClose={closePopup} onEndpoints={onEndpoints} />
        <FromLocalRepo open={popup === 'local'} onClose={closePopup} onEndpoints={onEndpoints} />
        <GettingStarted openPopup={openPopup} />
      </>);
  }

  return (
    <>
      <FromExample open={popup === 'example'} onClose={closePopup} onEndpoints={onEndpoints} />
      <FromGitHub open={popup === 'github'} onClose={closePopup} onEndpoints={onEndpoints} />
      <FromLocalRepo open={popup === 'local'} onClose={closePopup} onEndpoints={onEndpoints} />
      <AddEndpoints open={popup === 'endpoints'} onClose={closePopup} endpoints={endpoints} />
      <AddModal open={popup === 'add'} onClose={closePopup} openPopup={openPopup} />

      <Grid container sx={{ paddingBottom: 2, borderBottom: 1, borderColor: 'divider' }}>
        <Grid item container xs={6} alignItems="center">
          <ExtensionsStatus status={'ready'} />
        </Grid>
        <Grid item xs={6} alignItems="right" align="right">
          <Button color="secondary" variant="contained" onClick={openPopup('add')}>
            Add an Extension
          </Button>
        </Grid>
      </Grid>
      <Grid container spacing={2}>
        <Grid item xs={3} marginTop={1}>
          <MenuList sx={{ width: '100%' }}>
            {middlewares.map((m) => (
              <MenuItem
                key={m.slug}
                selected={m.slug === active}
                onClick={() => setActive(m.slug)}
              >
                <ExtensionStatus status={m.status} />
                {m.slug}
              </MenuItem>
            ))}
          </MenuList>
        </Grid>
        <Grid item xs={9}>
          <ExtensionSettings extension={middlewareBySlug(active)} reload={reload} />
        </Grid>
      </Grid>
    </>
  );
};

export default ExtensionsTab;
