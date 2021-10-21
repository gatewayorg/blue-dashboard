// import localeUS from './en-US.json';
import localeCN from './cn.json';

export type LocaleName = 'en-US' | "CN";

export interface ILocaleData {
  value: LocaleName;
  name: string;
}

// locale data
const locales: ILocaleData[] = [
  {
    name: 'English',
    value: 'en-US',
  },
  {
    name: '中文',
    value: 'CN',
  },
];

// locale data
const LOCALES_DATA = {
  'en-US': localeCN,
  'CN': localeCN
};

export { locales, LOCALES_DATA };
