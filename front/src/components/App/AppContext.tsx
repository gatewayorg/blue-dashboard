import React from 'react';

type Locale = 'en-US'|'CN';

export interface ILocaleContext {
  locale: Locale;
}

const LocaleContext = React.createContext<ILocaleContext>({ locale:'CN' });

export { LocaleContext };