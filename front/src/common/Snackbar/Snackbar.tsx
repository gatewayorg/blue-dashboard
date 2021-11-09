import React, { ReactNode } from 'react';
import { t } from '../utils/intl';
import { makeStyles } from '@material-ui/core';
import { Loading } from '../Loading/Loading';
import { Message, MessageType } from '../Loading/Message';

const useStyles = makeStyles(() => ({
  root: {
    height: '100%',
  }
}));

interface ISnackbarProps {
  loading: boolean;
  error: string;
  data?: any[];
  children: ReactNode;
}

const Snackbar = ({
  loading,
  error,
  data,
  children,
}: ISnackbarProps) => {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      {loading ? <Loading /> : 
        error ? <Message type={MessageType.Error} message={error} /> : 
          data?.length === 0 ? t('common.no-data') : children
      }
    </div>
  )
}

export { Snackbar };