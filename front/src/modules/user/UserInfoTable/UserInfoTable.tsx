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
import { UserInfoTableStyle } from "./UserInfoTableStyle";
import { IDataInfo } from '../userList';
import { authClient } from '../../../api/service';
import { NewSwitch } from "../../../common/Switch/NewSwitch"
import { useSnackbar } from 'notistack';

interface IAppsTableRowProps extends StyledComponentProps {
  theme: Theme;
  datadetail: IDataInfo;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setUserSelect: React.Dispatch<React.SetStateAction<IDataInfo | undefined>>;
  setPasswordStatus: React.Dispatch<React.SetStateAction<boolean>>;
  setUserStatus: React.Dispatch<React.SetStateAction<boolean>>;
  userNames: string;
  page: number;
  query: (page: number) => void
}

const UserInfoTableComponent = ({
  classes = {},
  theme,
  datadetail,
  userNames,
  setOpen,
  setUserSelect,
  query,
  page,
  setPasswordStatus,
  setUserStatus,
}: IAppsTableRowProps) => {
  const { enqueueSnackbar } = useSnackbar();
  const changeStatus = () => {
    setOpen(true);
    setUserSelect(datadetail);
  };
  const changePassword =()=>{
    setPasswordStatus(true);
    setUserSelect(datadetail);
  }
  const changeUser =()=>{
    setUserStatus(true);
    setUserSelect(datadetail);
  }

  const deleteRole = () => {
    authClient.delete(`/api/user/del?id=${datadetail.id}`)
      .then(res => {
        if (res.status === 200) {
          let timer;
          if(datadetail.name === window.sessionStorage.getItem('username')){
            enqueueSnackbar("Success", { 
              variant: 'success',
              autoHideDuration: 3000,
            })
            clearTimeout(timer);
            timer = setTimeout(()=>{
              window.sessionStorage.setItem('token',"");
              window.sessionStorage.setItem('username',"");
              window.location.href = '/login';
            },2000)
          }else{
            query(page);
          }
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  };

  const handleChange =(event:any)=>{
    authClient.patch(`/api/user/status`,{
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
        <TableCell>{datadetail.create_at}</TableCell>
        <TableCell>{datadetail.username}</TableCell>
        <TableCell>{datadetail.role_name}</TableCell>
        <TableCell>
          <Button onClick={deleteRole} className={classes.buttons}>
            {t("user-info.header.delete")}
          </Button>
          <Button onClick={changeStatus} className={classes.updates}>
            {t("user-info.header.change")}
          </Button>
          <Button onClick={changeUser} className={classes.updates}>
            {t("user-info.header.user")}
          </Button>
          {(userNames === datadetail.name) && <Button onClick={changePassword} className={classes.updates}>
            {t("user-info.header.password")}
          </Button>}
        </TableCell>
      </TableRow>
    </>
  );
};

const UserInfoTable = withStyles(UserInfoTableStyle)(
  withTheme(UserInfoTableComponent)
);

export { UserInfoTable };
