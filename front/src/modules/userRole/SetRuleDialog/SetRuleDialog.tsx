import React, { useEffect } from "react";
import { t } from '../../../common/utils/intl';
import { withStyles, Theme, Box, Button } from "@material-ui/core";
import { StyleRules } from "@material-ui/styles";
import { StyledComponentProps } from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import DialogContent from "@material-ui/core/DialogContent";
import DialogTitle from "@material-ui/core/DialogTitle";
import { onFormSubmit } from "../../../common/formTypes";
import Select from "@material-ui/core/Select";
import Input from "@material-ui/core/Input";
import MenuItem from "@material-ui/core/MenuItem";
import { Form, FormRenderProps } from "react-final-form";
// import { IApiRouters, IApiUserRouters } from "store/apiMappings/userMappings";
import FormControl from "@material-ui/core/FormControl";
import { IRoleList } from "../userRole"

const styles = (theme: Theme): StyleRules => ({
  root: {
    "& .MuiListItem-root.Mui-selected": {
      backgroundColor: "rgba(0, 0, 0, 0.4)"
    },
    "& .MuiInputBase-root": {
      margin: "0",
      border:'none',
    },
    "& .MuiSelect-nativeInput":{
      color:"#fff"
    }
  },
  selectControl: {
    width: "100%"
  },
  formControl: {
    width: "100%",
    border: "1px solid #ced4da",
    "& .MuiInput-underline:before": {
      borderBottom: "none"
    },
    "& .MuiSelect-select": {
      padding: '0 20px 0 10px',
      height:38,
      lineHeight:'38px',
      backgroundColor:'#191c23',
      color:'#fff',
    }
  },buttons: {
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

export interface IApiSetRoleValue {
  roleId: string;
}

export interface IAPICodeType {
  label: string;
  value: string;
}



interface IAppsFilterProps extends StyledComponentProps {
  openDialoy: boolean;
  selectedRoles?: string[];
  roleList?: IRoleList[];
  setSelectedRoles: React.Dispatch<React.SetStateAction<string[]>>;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onSubmit: (value: IApiSetRoleValue) => void;
}

const SetRuleDialogComponent = ({
  classes = {},
  selectedRoles,
  roleList,
  openDialoy,
  setOpen,
  setSelectedRoles,
  onSubmit
}: IAppsFilterProps) => {
  const [selectVal, setSelectVal] = React.useState<string[]>([]);

  const closePrintDialog = () => {
    setOpen(false);
  };

  useEffect(() => {
    if (selectedRoles && selectedRoles.length > 0) {
      const newRoleList = selectedRoles.map(items => {
        return items + "";
      });
      setSelectedRoles(newRoleList);
      setSelectVal(newRoleList);
    } else {
      setSelectVal([]);
      setSelectedRoles([]);
    }
  }, [selectedRoles,setSelectedRoles]);
  const changeSelect = (e: any) => {
    
    setSelectVal(e.target.value);
    setSelectedRoles(e.target.value);

  };
  const userForm = ({ handleSubmit }: FormRenderProps) => {
    return (
      <form onSubmit={handleSubmit} className={classes.root}>
        <Box m={-2} mt={2}>
          <Box p={2}>
            <Box pb={1}>{t("user-role.dialog.role")}</Box>
            <Box className={classes.selectControl}>
              <FormControl className={classes.formControl}>
                <Select
                  labelId="demo-mutiple-name-label"
                  id="demo-mutiple-name"
                  multiple
                  value={selectVal}
                  variant="outlined"
                  onChange={changeSelect}
                  input={<Input />}
                >
                  {roleList && roleList.map(items => (
                    <MenuItem key={items.id} value={items.id.toString()}>
                      {items.service + " - " + items.method}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Box>
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
              <Button onClick={closePrintDialog} color="primary" className={classes.buttons}>
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
        <span>{t("user-role.dialog.set")}</span>
      </DialogTitle>
      <DialogContent>
        <Form render={userForm} onSubmit={onSubmit as onFormSubmit} />
      </DialogContent>
    </Dialog>
  );
};
const SetRuleDialog = withStyles(styles)(SetRuleDialogComponent);

export { SetRuleDialog };
