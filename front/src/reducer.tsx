import React, { useReducer } from 'react';

export interface IState  {
  playIndex: number;
}

export const ActionTypes = {
  SET_PLAY_INDEX: 'SET_PLAY_INDEX',
}

const initState: IState = {
  playIndex: 0,
}

const AppContext = React.createContext<any>(initState);

const reducer = (state: IState, action: any) => {
  switch(action.type) {
    case ActionTypes.SET_PLAY_INDEX: 
      return {...state, playIndex: action.playIndex}
    default: return state;
  }
} 

const ContextProvider = (props:any) => {
  const [state, dispatch] = useReducer(reducer, initState)
  return (
    <AppContext.Provider value={{state, dispatch}}>
      {props.children}
    </AppContext.Provider>
  )
}

export { AppContext, reducer, ContextProvider }