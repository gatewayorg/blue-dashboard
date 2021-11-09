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
import { PasswordInputField } from "../../../common/Input/PasswordInputField"



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

export interface IChangePasswordValue {
  oldPassword: string;
  newPassword:string;
}

interface IAppsFilterProps extends StyledComponentProps {
  openDialoy: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onSubmit: (value: IChangePasswordValue) => void;
  initialValues?: Partial<IChangePasswordValue>;
}

const ChangePasswordDialogComponent = ({
  classes = {},
  openDialoy,
  setOpen,
  onSubmit,
  initialValues = {
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
            <Box pb={1}>{t("user-info.dialog.changePassword.oldPassword")}</Box>
            <Field name="oldPassword" component={PasswordInputField} margin="none" />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("user-info.dialog.changePassword.newPassword")}</Box>
            <Field name="newPassword" component={PasswordInputField} margin="none" />
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
        <span>{t("user-info.dialog.changePassword.title")}</span>
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
const ChangePasswordDialog = withStyles(styles)(
  ChangePasswordDialogComponent
);

export { ChangePasswordDialog };
