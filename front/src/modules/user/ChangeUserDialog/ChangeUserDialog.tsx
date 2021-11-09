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
import { InputField } from "../../../common/Input";

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

export interface IChangeUserValue {
  roleId: string;
  name:string;
  enable:string;
}

interface IAppsFilterProps extends StyledComponentProps {
  openDialoy: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onSubmit: (value: IChangeUserValue) => void;
  roleList?:IRoleInfo[];
  initialValues?: Partial<IChangeUserValue>;
  userSelect: IDataInfo | undefined
}
export const statusList = [
  { label: "true", value: "true" },
  { label: "false", value: "false" }
];

const ChangeUserDialogComponent = ({
  classes = {},
  openDialoy,
  setOpen,
  onSubmit,
  roleList,
  userSelect,
  initialValues = {
    roleId: userSelect?.role_id||"",
    name:userSelect?.name||"",
    enable:userSelect?.enable.toString()||statusList[0].value,
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
            <Box pb={1}>{t("user-info.dialog.addUser.name")}</Box>
            <Field name="name" component={InputField} margin="none" />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("user-info.dialog.addUser.role")}</Box>
            <Field
              name="roleId"
              component={SelectFormInput}
              margin="none"
              values={formatRoleList}
            />
          </Box>
          <Box p={2}>
            <Box pb={1}>{t("user-info.dialog.addUser.enable")}</Box>
            <Field
              name="enable"
              component={SelectFormInput}
              margin="none"
              values={statusList}
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
        <span>{t("user-info.dialog.addUser.title")}</span>
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
const ChangeUserDialog = withStyles(styles)(
  ChangeUserDialogComponent
);

export { ChangeUserDialog };
