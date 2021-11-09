import React, { Ref } from "react";
import TextField, { OutlinedTextFieldProps } from "@material-ui/core/TextField";
import { memo, useMemo } from "react";
import MenuItem, { MenuItemProps } from "@material-ui/core/MenuItem";

const MenuItemMemoized = memo(
  React.forwardRef((props: Omit<MenuItemProps, "button">, ref: Ref<any>) => {
    return <MenuItem {...props} innerRef={ref} />;
  }),
  (prev, next) => prev.value === next.value && prev.children === next.children
);

export interface ISelectOption {
  value?: string | number;
  label?: string;
  hideSelection?: boolean;
}

interface ISelectComponent extends OutlinedTextFieldProps {
  label?: string;
  values: ISelectOption[];
  placeholder?: string;
  defaultValue?: string;
}

const SHRINK = { shrink: true };
const SELECT_PROPS = {
  MenuProps: {
    keepMounted: true,
    MenuListProps: { style: { width: "100%" } }
  }
};

const Select = (props: ISelectComponent) => {
  const {
    label,
    values,
    value = undefined,
    fullWidth = true,
    variant = "outlined",
    defaultValue = "",
    ...rest
  } = props;

  const items = useMemo(() => {
    return values.map(option => {
      if (option?.hideSelection && value !== option.value) {
        return null;
      }
      return (
        <MenuItemMemoized
          key={option.value || ""}
          value={option.value || defaultValue}
        >
          {option.label}
        </MenuItemMemoized>
      );
    });
  }, [values, value, defaultValue]);

  return (
    <TextField
      variant={variant}
      select={true}
      margin="dense"
      label={label}
      InputLabelProps={SHRINK}
      value={value || defaultValue}
      fullWidth={fullWidth}
      SelectProps={SELECT_PROPS}
      {...rest}
    >
      {items}
    </TextField>
  );
};

export { Select };
