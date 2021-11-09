import React from 'react';
import { 
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';
import { PATH } from './const';
import { Home } from '../modules/index/home';
import { UserList } from '../modules/user/userList';
import { UserRule } from "../modules/userRule/userRule";
import { UserRole } from "../modules/userRole/userRole";
import { Login } from "../modules/login/login"
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
        <Route exact path={PATH.MY.My}>
          <UserList />
        </Route>
        <Route exact path={PATH.MY.Rule}>
          <UserRule />
        </Route>
        <Route exact path={PATH.MY.Role}>
          <UserRole />
        </Route>
        <Route exact path={PATH.LOGIN}>
          <Login />
        </Route>
        <Route path={''}>
          {/* <PageNotFound /> */}
        </Route>
      </Switch>
    </Router>
  )
}

export { MainRoutes } 