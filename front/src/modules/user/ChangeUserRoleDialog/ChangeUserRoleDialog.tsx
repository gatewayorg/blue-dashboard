import React,{useEffect} from "react";
import { t } from '../../../common/utils/intl';
import { withStyles, Theme, Box, Button } from "@material-ui/core";
import { StyleRules } from "@material-ui/styles";
import { StyledComponentProps } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { onFormSubmit } from "../../../common/formTypes";
import { Field, Form, FormRenderProps } from "react-final-form";
import { SelectFormInput } from "../../../common/Select";
import { IRoleInfo,IDataInfo } from "../userList";

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

export interface IUserStatusValue {
  role: string;
}

interface IAppsFilterProps extends StyledComponentProps {
  openDialoy: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onSubmit: (value: IUserStatusValue) => void;
  roleList?:IRoleInfo[];
  initialValues?: Partial<IUserStatusValue>;
  userSelect?:IDataInfo,
}

const ChangeUserRoleDialogComponent = ({
  classes = {},
  openDialoy,
  setOpen,
  onSubmit,
  roleList,
  userSelect,
  initialValues = {
    role: userSelect ? userSelect.role_id : ""
  }
}: IAppsFilterProps) => {
  const [formatRoleList, setFormatRoleList] = React.useState<any>();
  const closePrintDialog = () => {
    setOpen(false);
  };

  useEffect(() => {
    if (roleList) {
      const rolelists = roleList.map(items => {
        return {
          label: items.name,
          value: items.id
        };
      });
      setFormatRoleList(rolelists);
    }
  }, [roleList]);


  const userForm = ({ handleSubmit }: FormRenderProps) => {
    return (
      <form onSubmit={handleSubmit}>
        <Box m={-2} mt={2}  className={classes.root}>
          <Box p={2}>
            <Box pb={1}>{t("user-info.dialog.changeRole.role")}</Box>
            <Field
              name="role"
              component={SelectFormInput}
              margin="none"
              values={formatRoleList}
            />
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
                // disabled={isRequestInProgress(fetchChangeUserStatus)}
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
        <span>{t("user-info.dialog.changeRole.title")}</span>
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
const ChangeUserRoleDialog = withStyles(styles)(
  ChangeUserRoleDialogComponent
);

export { ChangeUserRoleDialog };
