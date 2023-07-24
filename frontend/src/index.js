import React from 'react';
import ReactDOM from 'react-dom/client';
import posthog from 'posthog-js';
import './index.css';
import {
  createBrowserRouter,
  RouterProvider
} from 'react-router-dom';
import { createTheme, ThemeProvider } from '@mui/material';
import Routes from './routes';
import Theme from './theme';

// Apoxy tracks usage of Proximal to help us improve the product
// and justify the continued investment in maintaining this tool.
// No personal information is collected.
// You can learn more about Apoxy here: https://apoxy.dev
// or you can email us at hello@apoxy.dev with questions.
// You can opt-out of tracking by commenting out the following line.
posthog.init('phc_k3l3Nr2HeDJwQMsPlIg2By3Beta8c4wZijdDmbQGEtE', { api_host: 'https://e.apoxy.dev' });

const router = createBrowserRouter(Routes);
const theme = createTheme(Theme);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>
);
