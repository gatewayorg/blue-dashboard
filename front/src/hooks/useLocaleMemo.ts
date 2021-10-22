import React from 'react';
import { DependencyList, useContext, useMemo } from 'react';

type Locale = 'en-US' |'CN';

export interface IAppContext {
  locale: Locale;
}
const AppContext = React.createContext<IAppContext>({ locale: 'CN' });

function useLocaleMemo<T = any>(
  memoFn: () => T,
  deps: DependencyList | undefined,
) {
  const context = useContext(AppContext);
  return useMemo(memoFn, [...(deps || []), context.locale]);
}

export { useLocaleMemo };
