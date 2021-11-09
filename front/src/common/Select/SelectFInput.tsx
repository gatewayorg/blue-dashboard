import React from "react";
import { FieldRenderProps } from "react-final-form";
import { FormControlProps } from "@material-ui/core/FormControl";
import { getErrorText, hasError } from "../form";
import { ISelectOption, Select } from "../SelectBasic/Select";

interface ISelectFieldProps extends FieldRenderProps<any> {
  values: ISelectOption[];
  formControlProps?: FormControlProps;
  placeholder?: string;
}

const SelectFormInput = ({
  input: { name, onChange, value },
  values,
  meta,
  ...rest
}: ISelectFieldProps & any) => {
  const showError = hasError(meta);

  return (
    <Select
      {...rest}
      name={name}
      helperText={getErrorText(meta)}
      error={!!showError}
      onChange={onChange}
      value={value}
      values={values}
    />
  );
};

export { SelectFormInput };
