import React, { ReactChild } from  'react';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles
} from '@material-ui/core';
import { Header } from './Header';

const styles = (theme: Theme): StyleRules => ({
  root: {
    width: '100%',
    // maxWidth: theme.spacing(160),
    // padding: theme.spacing(0, 8),
    margin: '0 auto',
    '@media (max-width: 769px)':{
      padding: theme.spacing(0, 3, 4),
      overflowY:'scroll',
      overflowX:"hidden",
    }
  },
  content: {
    // marginTop: theme.spacing(8),
  }
});

interface ILayoutProps
extends StyledComponentProps {
  children: ReactChild;
}

const LayoutComponent = ({
  classes = {},
  children,
}:ILayoutProps) => {
  return (
    <>
      <Header />
      <div className={classes.root}>
        <div className={classes.content}>
          {children}
        </div>
      </div>
    </>
  )
}

const Layout = withStyles(styles)(LayoutComponent);

export { Layout }