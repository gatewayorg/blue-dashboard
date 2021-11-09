import React, { useMemo } from 'react';

import {
  TextField,
  InputAdornment,
  IconButton,
  BaseTextFieldProps,
} from '@material-ui/core';
import { OutlinedTextFieldProps } from '@material-ui/core/TextField';


interface IPasswordInputComponent extends OutlinedTextFieldProps {
  label?: string;
  size: BaseTextFieldProps['size'];
}

const SHRINK = { shrink: true };
const PasswordInput = (props: IPasswordInputComponent) => {
  const [hidden, setHidden] = React.useState(true);

  const handleHidden = React.useCallback(() => {
    setHidden(!hidden);
  }, [hidden]);

  return useMemo(() => {
    const passwordType = hidden ? 'password' : 'text';

    const InputProps = {
      ...props.InputProps,
      endAdornment: (
        <InputAdornment position="end">
          <IconButton
            size="small"
            aria-label="toggle password visibility"
            onClick={handleHidden}
          >
          </IconButton>
        </InputAdornment>
      ),
    };

    return (
      <TextField
        margin="dense"
        InputLabelProps={SHRINK}
        {...props}
        type={passwordType}
        InputProps={InputProps}
      />
    );
  }, [hidden, props, handleHidden]);
};

PasswordInput.defaultProps = {
  fullWidth: true,
  size: 'medium',
};

export { PasswordInput };
