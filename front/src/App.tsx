import React from 'react';
import { MainRoutes } from './nav/MainRoutes';
import AppBase from './components/App/AppBase';
import { LOCALES_DATA } from './locales/locales';
import { mainTheme } from './themes/mainTheme'
import './index.css';

function App() {

  return (
    <AppBase
      translations={LOCALES_DATA}
      theme={mainTheme}
    >
      <MainRoutes />
    </AppBase>
  );
}

export default App;
