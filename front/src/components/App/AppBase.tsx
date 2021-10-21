import React from 'react';
import intl from 'react-intl-universal';
import { CssBaseline, MuiThemeProvider, Theme } from '@material-ui/core';
import { LocaleContext } from './AppContext';
import { ContextProvider } from '../../reducer';

interface IAppProps {
  translations: {};
  theme: Theme;
}

interface IAppState {
  initDone: boolean;
  locale: 'en-US'|"CN" ;
}

class AppBase extends React.Component<IAppProps, IAppState> {
  constructor(props: any) {
    super(props);
    this.state = { initDone: false, locale: 'CN'}
  }

  public componentDidMount(): void {
    this.loadLocales();
  }


  public render() {
    return (
      <ContextProvider>
        <LocaleContext.Provider value={{locale: this.state.locale}}>
          <MuiThemeProvider theme={this.props.theme}>
            <CssBaseline />
            {this.state.initDone && this.props.children}
          </MuiThemeProvider>
        </LocaleContext.Provider>
      </ContextProvider>
    )
  }

  protected loadLocales = () => {
    intl
      .init({
        currentLocale: this.state.locale,
        locales: this.props.translations,
        warningHandler: message => {
          
        },
        commonLocaleDataUrls: {},
      })
      .then(() => {
        this.setState({ initDone: true });
      });
  };
}

export default AppBase;