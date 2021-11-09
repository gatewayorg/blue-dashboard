import React from "react";

import { BaseTextFieldProps, TextField } from "@material-ui/core";
import { OutlinedTextFieldProps } from "@material-ui/core/TextField";

export interface IInputComponent extends OutlinedTextFieldProps {
  label?: string;
  size: BaseTextFieldProps["size"];
}

const SHRINK = { shrink: true };
const Input = (props: IInputComponent) => {
  return <TextField margin="dense" InputLabelProps={SHRINK} {...props} />;
};

Input.defaultProps = {
  fullWidth: true,
  size: "medium" as "medium"
};

export { Input };
