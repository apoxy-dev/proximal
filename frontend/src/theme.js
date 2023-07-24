const Theme = {
  palette: {
    primary: {
      main: '#1F1C1C',
    },
    secondary: {
      main: '#48CD84',
    },
  },
  typography: {
    fontFamily: '"IBM Plex Sans", sans-serif',
    h1: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    h2: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    h3: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    h4: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    h5: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    h6: {
      fontFamily: '"IBM Plex Mono", monospace',
    },
    button: {
      fontWeight: 'bold',
      textTransform: 'none',
    },
  },
  components: {
    MuiTab: {
      styleOverrides: {
        root: {
          fontSize: '1rem',
          textTransform: 'uppercase',
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
          borderRadius: '3px',
          '&:hover': {
            backgroundColor: '#1F1C1C',
            color: '#48CD84',
          },
        },
      },
    },
    MuiTableHead: {
      styleOverrides: {
        root: {
          '& th': {
            fontWeight: 'bold',
          },
        },
      },
    }
  },
};

export default Theme;
