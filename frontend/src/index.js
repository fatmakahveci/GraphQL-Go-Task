import React from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import App from './App';
import Graphi from './Graphi';

createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
    <Graphi />
  </React.StrictMode>
);
