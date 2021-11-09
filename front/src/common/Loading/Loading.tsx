import React from 'react';
import { t } from '../../common/utils/intl';
import url from '../../assets/icons/loading.gif';
import { StyledComponentProps, StyleRules, Theme, withStyles } from '@material-ui/core';

const styles = (theme: Theme): StyleRules => ({
  root: {
    display: 'flex',
    width: '100%',
    height: theme.spacing(12),
    alignItems: 'center',
    justifyContent: 'center',
    '& img': {
      width: theme.spacing(12),
    }
  }
});

interface ILoadingProps
extends StyledComponentProps {

}

const LoadingComponent = ({
  classes = {},
}: ILoadingProps) => {
  return (
    <div className={classes.root}>
      <img src={url} alt={t('common.loading')} />
    </div>
  )
}

const Loading = withStyles(styles)(LoadingComponent);

export { Loading }