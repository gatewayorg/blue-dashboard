import React,{useMemo,useCallback} from 'react';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles,
  Box,
  Button
} from '@material-ui/core';
import { Layout } from '../../components/Layout/Layout';
import { authClient } from '../../api/service';
import { useInitEffect } from '../../hooks/useInitEffect';
import { UserInfoTable } from './UserInfoTable/UserInfoTable'
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Table from "@material-ui/core/Table";
import { t } from '../../common/utils/intl';
import { ChangeUserRoleDialog } from './ChangeUserRoleDialog/ChangeUserRoleDialog';
import { AddUserDialog,IAddUserValue } from  './AddUserDialog/AddUserDialog';
import { ChangePasswordDialog,IChangePasswordValue } from './ChangePasswordDialog/ChangePasswordDialog'
import {ChangeUserDialog,IChangeUserValue} from './ChangeUserDialog/ChangeUserDialog'
import { useSnackbar } from 'notistack';
import { Pagination } from '../../common/Pagination'
import { ITEMS_PER_PAGE } from "../const"
import { useIsXSDown } from "../../hooks/themeHooks";

const styles = (theme: Theme): StyleRules => ({
  root:{
    padding:'20px',
    [theme.breakpoints.down("xs")]: {
      padding:'0'
    },
    '& .MuiTableCell-root':{
      [theme.breakpoints.down("xs")]: {
        padding:'10px'
      }
    },
  },
  buttons: {
    marginRight: 10,
    backgroundColor:'#099639',
    fontWeight: 500,
    lineHeight:'1.75',
    '&&:hover': {
      backgroundColor: '#099639',
      color:'#fff',
      opacity:'0.85'
    }
  },

});
export interface IDataInfo {
  create_at:string;
  enable:boolean;
  id:string;
  name:string;
  username:string;
  role_detail:string;
  role_id:string;
  role_name:string;
}
export interface IRoleInfo {
  create_time:string;
  detail:string;
  enable:boolean;
  id:string;
  name:string;
}

interface IHomeProps
extends StyledComponentProps {

}

const UserListComponent = ({
  classes = {},
}: IHomeProps) => {
  const [dataInfo, setDataInfo] = React.useState<IDataInfo[]>();
  const [roleList, setRoleList] = React.useState<IRoleInfo[]>();
  const [open, setOpen] = React.useState(false);
  const [userSelect, setUserSelect] = React.useState<IDataInfo>();
  const [addStatus, setAddStatus] = React.useState(false);
  const [passwordStatus, setPasswordStatus] = React.useState(false);
  const [userStatus, setUserStatus] = React.useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const [userNames, setUserNames] = React.useState("");
  const [total, setTotal] = React.useState(0);
  const [page, setPage] = React.useState(0);
  const isXSDown = useIsXSDown();

  const query= useCallback((page:number) => {
    authClient.get(`/api/user/list?page=${page+1}&page_size=${ITEMS_PER_PAGE}`)
      .then(res => {
        if (res.status === 200) {
          setDataInfo(res.data.data)
          setTotal(parseInt(res.data.total))
        }
      }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  },[enqueueSnackbar]);

  useInitEffect(() => {
    query(page);
    setUserNames(window.sessionStorage.getItem('username')||"")
  })

  useInitEffect(() => {
    authClient.get(`/api/rbac/role/list?page=1&page_size=10`)
      .then(res => {
        if (res.status === 200) {
          setRoleList(res.data.data);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  })

  const items = useMemo(() =>
    dataInfo && dataInfo.map(datadetail => {
        return (
          <UserInfoTable
          key={datadetail.id}
          datadetail={datadetail}
          setOpen={setOpen}
          setUserSelect={setUserSelect}
          setPasswordStatus={setPasswordStatus}
          setUserStatus={setUserStatus}
          query={query}
          page={page}
          userNames={userNames}
          />
        );
      }),
    [dataInfo, page, query, userNames]
  );

  const onchangeRole = (value: any) => {
    authClient.patch(`/api/user/select/role`,{
      "id": userSelect?.id,
      "role_id": value.role
    })
      .then(res => {
        if (res.status === 200) {
          setOpen(false);
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  };
  const onAddUser = (value:IAddUserValue)=>{
    authClient.post(`/api/user/add`,{
      "username": value.username,
      "name": value.name,
      "role_id":value.roleId,
      "enable":value.enable === "true"?true:false,
      "passwd":value.password
    })
      .then(res => {
        if (res.status === 200) {
          setAddStatus(false);
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }
  const addUser=()=>{
    setAddStatus(true);
  }
  const onChangePassword =(value:IChangePasswordValue)=>{
    authClient.patch(`/api/user/pwd`,{
      "new_passwd": value.newPassword,
      "old_passwd": value.oldPassword,
    })
      .then(res => {
        if (res.status === 200) {
          setPasswordStatus(false);
          window.sessionStorage.setItem('token',"");
          window.sessionStorage.setItem('username',"");
          window.location.href = '/login';
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }

  const onChangeUser = (value:IChangeUserValue)=>{
    authClient.put(`/api/user`,{
      "name": value.name,
      "role_id":value.roleId,
      "enable":value.enable === "true"?true:false,
      "id": userSelect?.id,
    })
      .then(res => {
        if (res.status === 200) {
          setUserStatus(false);
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }

  const handlePageChange = (page: number) => {
    setPage(page);
    query(page);
  };

  return (
    <Layout>
      <div className={classes.root}>
      <Box display="flex" justifyContent="end">
        <Button onClick={addUser} className={classes.buttons}>
          {t("user-info.header.add")}
        </Button>
      </Box>
      <Table>
          <TableHead>
            <TableRow>
            {!isXSDown && <TableCell>{t("user-info.header.id")}</TableCell>}
              <TableCell>{t("user-info.header.name")}</TableCell>
              <TableCell>{t("user-info.header.enable")}</TableCell>
              {!isXSDown && <TableCell>{t("user-info.header.create-time")}</TableCell>}
              {!isXSDown && <TableCell>{t("user-info.header.username")}</TableCell>}
              <TableCell>{t("user-info.header.role-name")}</TableCell>
              <TableCell>{t("user-info.header.operating")}</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>{items}</TableBody>
        </Table>
        {total > ITEMS_PER_PAGE && (
        <Box pt={2} display="flex" justifyContent="flex-end">
          <Pagination
            count={total}
            perPage={ITEMS_PER_PAGE}
            onPageChange={handlePageChange}
            page={page}
          />
        </Box>
      )}
        <ChangeUserRoleDialog
        openDialoy={open}
        setOpen={setOpen}
        onSubmit={onchangeRole}
        roleList={roleList}
        userSelect={userSelect}
      />
      <AddUserDialog 
        openDialoy={addStatus}
        setOpen={setAddStatus}
        onSubmit={onAddUser}
        roleList={roleList}
      />
      <ChangePasswordDialog 
        openDialoy={passwordStatus}
        setOpen={setPasswordStatus}
        onSubmit={onChangePassword}
      />
      <ChangeUserDialog 
        openDialoy={userStatus}
        setOpen={setUserStatus}
        onSubmit={onChangeUser}
        roleList={roleList}
        userSelect={userSelect}
      />
      </div>
    </Layout>
  )
}

const UserList = withStyles(styles)(UserListComponent);
export { UserList };