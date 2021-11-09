import React from "react";
import { t } from '../../../common/utils/intl';
import { withStyles, Theme, Box, Button } from "@material-ui/core";
import { StyleRules } from "@material-ui/styles";
import { StyledComponentProps } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { onFormSubmit } from "../../../common/formTypes";
import { Field, Form, FormRenderProps } from "react-final-form";
import { InputField } from "../../../common/Input";
import { SelectFormInput } from "../../../common/Select";
import { IDataInfo } from "../userRole"

export interface IAPICodeType {
  label: string;
  value: string;
}

const styles = (theme: Theme): StyleRules => ({
  root:{
    "& .MuiSelect-select": {
      backgroundColor:'#191c23',
      color:'#fff'
    },
    "& .MuiInputBase-root":{
      border:'none',
      margin:'8px 0'
    },
    "& .MuiOutlinedInput-input":{
      padding:0,
      backgroundColor:'#191c23',
      color:'#fff',
      textIndent:'16px'
    },
    "& .MuiBox-root":{
      paddingTop:0,
    },
    "& .MuiOutlinedInput-input:-webkit-autofill":{
      boxShadow:'0 0 0 100px #191c23 inset'
    }
  },

  buttons: {
    marginRight: 10,
    backgroundColor:'#099639',
    fontWeight: 500,
    lineHeight:'1.75',
    color:'#fff',
    '&&:hover': {
      backgroundColor: '#099639',
      color:'#fff',
      opacity:'0.85'
    }
  },
});

export interface IAddRoleValue {
  name: string;
  enable:string;
  detail:string;
}

interface IAppsFilterProps extends StyledComponentProps {
  openDialoy: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onSubmit: (value: IAddRoleValue) => void;
  initialValues?: Partial<IAddRoleValue>;
  selectedRole: IDataInfo | undefined;
}
export const statusList = [
  { label: "true", value: "true" },
  { label: "false", value: "false" }
];

const UpdateRoleDialogComponent = ({
  classes = {},
  openDialoy,
  setOpen,
  onSubmit,
  selectedRole,
  initialValues = {
    name:selectedRole?selectedRole.name:"",
    detail:selectedRole?selectedRole.detail:"",
    enable:selectedRole?selectedRole.enable.toString():"",
  }
}: IAppsFilterProps) => {

  const closePrintDialog = () => {
    setOpen(false);
  };




  const userForm = ({ handleSubmit }: FormRenderProps) => {
    return (
      <form onSubmit={handleSubmit}>
        <Box m={-2} mt={2}  className={classes.root}>
          <Box p={2}>
            <Box pb={1}>{t("user-role.dialog.name")}</Box>
            <Field name="name" component={InputField} margin="none" />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("user-role.dialog.enable")}</Box>
            <Field
              name="enable"
              component={SelectFormInput}
              margin="none"
              values={statusList}
            />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("user-role.dialog.detail")}</Box>
            <Field name="detail" component={InputField} margin="none" />
          </Box>
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            p={2}
            mb={2}
            flex="1"
          >
            <Box>
              <Button onClick={closePrintDialog} color="primary"  className={classes.buttons}>
                {t("user-info.dialog.changeRole.cancel")}
              </Button>
            </Box>
            <Box ml={3}>
              <Button
                type="submit"
                color="primary"
                className={classes.buttons}
              >
                {t("user-info.dialog.changeRole.submit")}
              </Button>
            </Box>
          </Box>
        </Box>
      </form>
    );
  };

  return (
    <Dialog
      open={openDialoy}
      onClose={closePrintDialog}
      fullWidth={true}
      aria-labelledby="form-dialog-title"
    >
      <DialogTitle id="form-dialog-title">
        <span>{t("user-role.dialog.update")}</span>
      </DialogTitle>
      <DialogContent>
        <Form
          render={userForm}
          onSubmit={onSubmit as onFormSubmit}
          initialValues={initialValues}
        />
      </DialogContent>
    </Dialog>
  );
};
const UpdateRoleDialog = withStyles(styles)(
  UpdateRoleDialogComponent
);

export { UpdateRoleDialog };
