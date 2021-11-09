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
import { UserRoleTable } from './UserRoleTable/UserRoleTableTable'
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Table from "@material-ui/core/Table";
import { t } from '../../common/utils/intl';
import{ AddRoleDialog,IAddRoleValue } from './AddRoleDialog/AddRoleDialog'
import{ UpdateRoleDialog } from './UpdateRoleDialog/UpdateRoleDialog'
import { SetRuleDialog } from './SetRuleDialog/SetRuleDialog'
import { useSnackbar } from 'notistack';
import { Pagination } from '../../common/Pagination'
import { ITEMS_PER_PAGE } from "../const"

const styles = (theme: Theme): StyleRules => ({
  root:{
    padding:'20px'
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
  id:string;
  name:string;
  enable:boolean;
  detail:string;
  create_time:string;
  rule_ids:string[];
}

export interface IRoleList{
  detail: string;
  id:  string;
  method: string;
  service: string;
}

interface IHomeProps
extends StyledComponentProps {

}

const UserRoleComponent = ({
  classes = {},
}: IHomeProps) => {
  const [dataInfo, setDataInfo] = React.useState<IDataInfo[]>();
  const [roleList, setRoleList] = React.useState<IRoleList[]>();
  const [open, setOpen] = React.useState(false);
  const [updateStatus, setUpdateStatus] = React.useState(false);
  const [ruleStatus, setTuleStatus] = React.useState(false);
  const [selectedRole, setSelectedRole] = React.useState<IDataInfo>();
  const [selectedRoles, setSelectedRoles] = React.useState<string[]>([""]);
  const [total, setTotal] = React.useState(0);
  const [page, setPage] = React.useState(0);

  const { enqueueSnackbar } = useSnackbar();
  const query= useCallback((page:number) => {
    authClient.get(`/api/rbac/role/list?page=${page+1}&page_size=${ITEMS_PER_PAGE}`)
      .then(res => {
        if (res.status === 200) {
          setDataInfo(res.data.data);
          setTotal(parseInt(res.data.total));
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
    authClient.get(`/api/rbac/rule/list?page=1&page_size=1000`)
      .then(res => {
        if (res.status === 200) {
          setRoleList(res.data.data)
        }
      }).catch(err => {
        enqueueSnackbar(err.response.data.message||"Error", { 
          variant: 'error',
          autoHideDuration: 3000,
        })
      })
  })

  const addRole  = ()=>{
    setOpen(true);
  }
  const  onAddRole = (value:IAddRoleValue)=>{
    authClient.post(`/api/rbac/role`,{
        "name": value.name,
        "enable": value.enable ==="ture"?true:false,
        "detail":value.detail
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
  }
  const onUpdatedRole=(value:any)=>{
    authClient.put(`/api/rbac/role`,{
      "name": value.name,
      "enable": value.enable ==="true"?true:false,
      "detail":value.detail,
      "id":selectedRole?.id
    })
      .then(res => {
        if (res.status === 200) {
          setUpdateStatus(false);
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }

  const onSetRule = ()=>{
    const newData  = selectedRoles[0] === "0"? selectedRoles?.slice(1,).toString() :selectedRoles.toString();
    authClient.patch(`/api/user/select/role`,{
      "role_id": newData,
      "id":selectedRole?.id
    })
      .then(res => {
        if (res.status === 200) {
          setUpdateStatus(false);
          query(page);
        }
    }).catch(err => {
      enqueueSnackbar(err.response.data.message||"Error", { 
        variant: 'error',
        autoHideDuration: 3000,
      })
    })
  }

  const items = useMemo(() =>
    dataInfo && dataInfo.map(datadetail => {
        return (
          <UserRoleTable
          key={datadetail.id}
          datadetail={datadetail}
          query={query}
          setUpdateStatus={setUpdateStatus}
          setSelectedRole={setSelectedRole}
          setTuleStatus={setTuleStatus}
          page={page}
          />
        );
      }),
    [dataInfo, page, query]
  );

  const handlePageChange = (page: number) => {
    setPage(page);
    query(page);
  };

  return (
    <Layout>
      <div className={classes.root}>
      <Box display="flex" justifyContent="end">
        <Button onClick={addRole} className={classes.buttons}>
          {t("user-role.header.add")}
        </Button>
      </Box>
      <Table>
          <TableHead>
            <TableRow>
              <TableCell>{t("user-role.header.id")}</TableCell>
              <TableCell>{t("user-role.header.name")}</TableCell>
              <TableCell>{t("user-role.header.enable")}</TableCell>
              <TableCell>{t("user-role.header.create-time")}</TableCell>
              <TableCell>{t("user-role.header.detail")}</TableCell>
              <TableCell>{t("user-rule.header.operating")}</TableCell>
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
        <AddRoleDialog 
          openDialoy={open}
          setOpen={setOpen}
          onSubmit={onAddRole}
        />
        <UpdateRoleDialog 
          openDialoy={updateStatus}
          setOpen={setUpdateStatus}
          onSubmit={onUpdatedRole}
          selectedRole={selectedRole}
        />
        <SetRuleDialog
          openDialoy={ruleStatus}
          setOpen={setTuleStatus}
          onSubmit={onSetRule} 
          selectedRoles={selectedRole?.rule_ids} 
          setSelectedRoles={setSelectedRoles}
          roleList={roleList}
        />
      </div>
      
    </Layout>
  )
}

const UserRole = withStyles(styles)(UserRoleComponent);
export { UserRole };