import React from 'react';
import { 
  StyledComponentProps,
  Theme, 
  Typography,
  withStyles
} from '@material-ui/core';
import classNames from 'classnames';

const styles = (theme: Theme) => ({
  error: {
    color: theme.palette.error.main,
  },
  success: {
    color: theme.palette.success.main,
  }
});

export enum MessageType {
  Error = 'error',
  Success = 'success',
}

interface IMessageProps
extends StyledComponentProps {
  message: string;
  type: MessageType,
}

const MessageComponent = ({
  classes = {},
  message,
  type,
}: IMessageProps) => {
  return <Typography variant='body1' className={classNames(classes.root, type === MessageType.Error && classes.error, type === MessageType.Success && classes.success)}>{message}</Typography>
}

const Message = withStyles(styles)(MessageComponent);

export { Message }