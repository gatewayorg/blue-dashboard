import React from "react";
import {
  StyledComponentProps,
  Theme,
  Button
} from "@material-ui/core";
import { t } from '../../../common/utils/intl';
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import withTheme from "@material-ui/core/styles/withTheme";
import withStyles from "@material-ui/core/styles/withStyles";
import { UserInfoTableStyle } from "./UserRoleTableStyle";
import { IDataInfo } from '../userRole';
import { authClient } from '../../../api/service';
import { NewSwitch } from "../../../common/Switch/NewSwitch"
import { useSnackbar } from 'notistack';

interface IAppsTableRowProps extends StyledComponentProps {
  theme: Theme;
  datadetail: IDataInfo;
  query: (page: number) => void;
  setUpdateStatus: React.Dispatch<React.SetStateAction<boolean>>;
  setTuleStatus: React.Dispatch<React.SetStateAction<boolean>>;
  setSelectedRole: React.Dispatch<React.SetStateAction<IDataInfo | undefined>>
  page: number;
}

const UserRoleTableComponent = ({
  classes = {},
  theme,
  datadetail,
  page,
  query,
  setSelectedRole,
  setUpdateStatus,
  setTuleStatus,
}: IAppsTableRowProps) => {
  const { enqueueSnackbar } = useSnackbar();
  const deleteRole = () => {
    authClient.delete(`/api/rbac/role?id=${datadetail.id}`)
      .then(res => {
        if (res.status === 200) {
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  };
  const handleChange =(event:any)=>{
    authClient.patch(`/api/rbac/role/status`,{
      "enable": event.target.checked,
      "id":datadetail.id
    })
      .then(res => {
        if (res.status === 200) {
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }
  const updateRole =()=>{
    setUpdateStatus(true);
    setSelectedRole(datadetail);
  }

  const setRule =()=>{
    setTuleStatus(true);
    setSelectedRole(datadetail);
  }

  return (
    <>
      <TableRow
        key={datadetail.id}
        hover
        style={{ cursor: "pointer" }}
      >
        <TableCell>{datadetail.id}</TableCell>
        <TableCell>{datadetail.name}</TableCell>
        <TableCell><NewSwitch checked={datadetail.enable} onChange={handleChange}/></TableCell>
        <TableCell>{datadetail.create_time}</TableCell>
        <TableCell>{datadetail.detail}</TableCell>
        <TableCell>
          <Button onClick={deleteRole} className={classes.buttons}>
            {t("user-role.header.delete")}
          </Button>
          <Button onClick={updateRole} className={classes.updates}>
            {t("user-role.header.update")}
          </Button>
          <Button onClick={setRule} className={classes.updates}>
            {t("user-role.header.set")}
          </Button>
        </TableCell>
      </TableRow>
    </>
  );
};

const UserRoleTable = withStyles(UserInfoTableStyle)(
  withTheme(UserRoleTableComponent)
);

export { UserRoleTable };
