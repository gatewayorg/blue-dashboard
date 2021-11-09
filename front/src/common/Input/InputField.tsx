import React from "react";
import { FieldRenderProps } from "react-final-form";
import { Input } from "./Input";
import { getErrorText, hasError } from "../form";
import { OutlinedInputProps } from "@material-ui/core/OutlinedInput";

interface IInputFieldProps extends FieldRenderProps<HTMLInputElement> {
  showErrorText?: boolean;
  showErrorOnSubmitOnly?: boolean;
  InputProps?: Partial<OutlinedInputProps>;
}

const InputField = ({
  input: { name, onChange, value, onBlur, onFocus },
  meta,
  showErrorText = true,
  showErrorOnSubmitOnly = false,
  ...rest
}: IInputFieldProps) => {
  const showError = showErrorOnSubmitOnly
    ? meta.submitFailed && hasError(meta)
    : hasError(meta);

  const customOnChange = (event: any) => {
    onChange(event);
  };

  return (
    <Input
      {...rest}
      variant="outlined"
      name={name}
      helperText={showErrorText && getErrorText(meta)}
      error={!!showError}
      onChange={customOnChange}
      // @ts-ignore
      onBlur={onBlur}
      // @ts-ignore
      onFocus={onFocus}
      value={value}
    />
  );
};

export { InputField };
