import React,{useMemo,useCallback} from 'react';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles,
  Box
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
import{  SetRuleDialog,ISetRuleValue } from './SetRuleDialog/SetRuleDialog'
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
  method:string;
  detail:string;
  service:string;
}

interface IHomeProps
extends StyledComponentProps {

}

const UserRuleComponent = ({
  classes = {},
}: IHomeProps) => {
  const [dataInfo, setDataInfo] = React.useState<IDataInfo[]>();
  const [open, setOpen] = React.useState(false);
  const [userSelect, setUserSelect] = React.useState<IDataInfo>();
  const { enqueueSnackbar } = useSnackbar();
  const [total, setTotal] = React.useState(0);
  const [page, setPage] = React.useState(0);

  const query= useCallback((page:number) => {
    authClient.get(`/api/rbac/rule/list?page=${page+1}&page_size=${ITEMS_PER_PAGE}`)
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
  })

  const onSetRule =(value:ISetRuleValue)=>{
    authClient.patch(`/api/rbac/set/detail`,{
      "id":userSelect?.id,
      "detail": value.detail,
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
  const handlePageChange = (page: number) => {
    setPage(page);
    query(page);
  };
  
  const items = useMemo(() =>
    dataInfo && dataInfo.map(datadetail => {
        return (
          <UserRoleTable
          key={datadetail.id}
          datadetail={datadetail}
          setUserSelect={setUserSelect}
          setOpen={setOpen}
          />
        );
      }),
    [dataInfo]
  );

  return (
    <Layout>
      <div className={classes.root}>
      <Table>
          <TableHead>
            <TableRow>
              <TableCell>{t("user-rule.header.id")}</TableCell>
              <TableCell>{t("user-rule.header.method")}</TableCell>
              <TableCell>{t("user-rule.header.service")}</TableCell>
              <TableCell>{t("user-rule.header.detail")}</TableCell>
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
        <SetRuleDialog 
          openDialoy={open}
          setOpen={setOpen}
          onSubmit={onSetRule}
        />
      </div>
      
    </Layout>
  )
}

const UserRule = withStyles(styles)(UserRuleComponent);
export { UserRule };