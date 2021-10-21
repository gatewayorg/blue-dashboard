import React from 'react';
import { 
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';
import { PATH } from './const';
import { Home } from '../modules/index/home'
interface IMainNavigationProps {
    
}

function MainRoutes(_: IMainNavigationProps) {



  React.useEffect(() => {

  }, []);

  return (
    <Router>
      <Switch>
        <Route exact path={PATH.DASHBOARD}>
          <Home />
        </Route>
        <Route path={''}>
          {/* <PageNotFound /> */}
        </Route>
      </Switch>
    </Router>
  )
}

export { MainRoutes } 