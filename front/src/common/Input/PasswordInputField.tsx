import React from 'react';
import { FieldRenderProps } from 'react-final-form';
import { PasswordInput } from './PasswordInput';
import { getErrorText, hasError } from '../form';

interface IPasswordInputFieldProps extends FieldRenderProps<HTMLInputElement> {
  showErrorText?: boolean;
  showErrorOnSubmitOnly?: boolean;
}

const PasswordInputField = ({
  input: { name, onChange, value, onBlur, onFocus },
  meta,
  showErrorText = true,
  showErrorOnSubmitOnly = false,
  ...rest
}: IPasswordInputFieldProps) => {
  const showError = showErrorOnSubmitOnly
    ? meta.submitFailed && hasError(meta)
    : hasError(meta);

  return (
    <PasswordInput
      {...rest}
      variant="outlined"
      name={name}
      helperText={showErrorText && getErrorText(meta)}
      error={!!showError}
      onChange={onChange}
      // @ts-ignore
      onBlur={onBlur}
      // @ts-ignore
      onFocus={onFocus}
      value={value}
    />
  );
};

export { PasswordInputField };
