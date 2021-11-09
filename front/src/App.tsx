import React from 'react';
import { MainRoutes } from './nav/MainRoutes';
import AppBase from './components/App/AppBase';
import { LOCALES_DATA } from './locales/locales';
import { mainTheme } from './themes/mainTheme'
import './index.css';
import { SnackbarProvider } from 'notistack';





function App() {

  return (
    <AppBase
      translations={LOCALES_DATA}
      theme={mainTheme}
    >
      <SnackbarProvider maxSnack={1}>
      <MainRoutes />
      </SnackbarProvider>
    </AppBase>
  );
}

export default App;
