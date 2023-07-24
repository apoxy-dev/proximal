import { Navigate } from 'react-router-dom';

import Main from './templates/Main';

import NotFound from './pages/NotFound';
import Home from './pages/home';

const Routes = [
  {
    path: '/',
    element: <Main />,
    errorElement: <NotFound />,
    children: [
      {
        index: true,
        element: <Navigate to="/extensions" />,
      },
      {
        path: '/extensions',
        element: <Home />,
      },
      {
        path: '/endpoints',
        element: <Home />,
      },
      {
        path: '/logs',
        element: <Home />,
      },
    ]
  },
];

export default Routes;
