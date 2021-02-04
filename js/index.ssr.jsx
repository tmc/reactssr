import * as React from 'react';
import * as Server from 'react-dom/server'
import './index.css';
import App from './App';

const AppOutput = Server.renderToString(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// reactssr.render is the callback injected by the go runtime to pass the rendered application back.
reactssr.render(AppOutput);

