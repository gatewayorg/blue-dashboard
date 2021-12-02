import React from 'react';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles
} from '@material-ui/core';
import { Layout } from '../../components/Layout/Layout';
import { StatusCodeCount } from "./statusCodeCountChart";
import { RequestTotalCharts } from "./requestTotalCharts"
import { authClient } from '../../api/service';
import { useInitEffect } from '../../hooks/useInitEffect';
import { CacheTotalCharts } from "./cacheTotalCharts";
import { SharedMemoryChart } from "./sharedMemoryChart"
import { useSnackbar } from 'notistack';

const styles = (theme: Theme): StyleRules => ({
  root:{
    
  },
  block: {
    display:"flex",
    margin:"20px 20px",
    borderRadius:"10px",
    overflow:'hidden',
    [theme.breakpoints.down("xs")]: {
      flexFlow:"column",
      margin:"0px",
      marginTop:"10px",
    }
  },
  errorCount:{
    backgroundColor:'#282c34',
    marginLeft:'10px',
    borderRadius:'8px',
    width:"15%",
    height:300,
    textAlign:'center',
    border: '1px solid rgba(256,256,256,0.1)',
    '&:first-child': {
      marginLeft:'0'
    },
    [theme.breakpoints.down("xs")]: {
      width:"100%",
      marginLeft:'0',
      marginTop:'10px',
      height:100,
    }
  },
  requestCount:{
    backgroundColor:'#282c34',
    marginLeft:'10px',
    borderRadius:'8px',
    width:"20%",
    height:300,
    textAlign:'center',
    border: '1px solid rgba(256,256,256,0.1)',
    [theme.breakpoints.down("xs")]: {
      width:"100%",
      marginLeft:'0',
      marginTop:'10px',
    }
  },
  countTitle:{
    fontWeight:600,
    fontSize:16,
    backgroundColor:'#282c34',
    margin:0,
    textAlign:'center',
    lineHeight:'48px',
    borderTopLeftRadius:'8px',
    borderTopRightRadius:'8px'
  },
  count:{
    color:'#099639',
    fontSize:"30px",
    textAlign:'center',
    margin:0,
    fontWeight:500,
    lineHeight:'200px',
    [theme.breakpoints.down("xs")]: {
      lineHeight:'40px',
    }
  },
  statusCodeCount:{
    width:"calc(100% - 20px)",
    border: '1px solid rgba(256,256,256,0.2)',
    borderRadius:"8px",
    [theme.breakpoints.down("xs")]: {
      width:"100%",
    }
  },
  requestTotalContent:{
    height:250,
    overflow:'hidden',
    borderBottomLeftRadius:'8px',
    borderBottomRightRadius:'8px',
  },
  cacheTotal:{
    width: 'calc(50% - 50px)',
    marginLeft:"10px",
    overflow:"hidden",
    [theme.breakpoints.down("xs")]: {
      width: '100%',
      marginLeft:"0",
      marginTop:'10px',
    }
  },
  chartContent:{
    height:'252px'
  },
  hostInfo:{
    width:"15%",
    backgroundColor: '#282c34',
    borderRadius:'8px',
    marginLeft:'10px',
    border: '1px solid rgba(256,256,256,0.2)',
    '&:first-child': {
      marginLeft:'0'
    },
    [theme.breakpoints.down("xs")]: {
      width:"100%",
      marginLeft:'0',
      marginTop:'10px',
    }
  },
  hostDetail:{
    color:'#099639',
    fontSize:"22px",
    textAlign:'center',
    margin:0,
    fontWeight:500,
    lineHeight:'80px',
  }
});

export interface IDataInfo {
  cache_total: ICacheTotal;
  connections:IConnections[];
  host_name:string;
  ip:string;
  request_total:IRequestTotal;
  shared_memory:ISharedMemory[];
  status:string;
  version:string;
}

export interface ICacheTotal {
  bypass: string;
  expired:string;
  hit:string;
  miss:string;
  revalidated:string;
  scarce:string;
  stale:string;
  updating:string;
}
export interface IConnections {
  accepted: string;
  active:string;
  handled:string;
  reading:string;
  requests:string;
  time:string;
  waiting:string;
  writing:string;
}
export interface IRequestTotal {
  in_bytes: string;
  out_bytes:string;
  _1xx:string;
  _2xx:string;
  _3xx:string;
  _4xx:string;
  _5xx:string;
}
export interface ISharedMemory{
  max_size:string;
  time:string;
  used_node:string;
  used_size:string;
}

interface IHomeProps
extends StyledComponentProps {

}

const HomeComponent = ({
  classes = {},
}: IHomeProps) => {
  const [dataInfo, setDataInfo] = React.useState<IDataInfo[]>();
  const { enqueueSnackbar } = useSnackbar();

  useInitEffect(() => {
    authClient.get(`/api/index?start=${parseInt((new Date().getTime()/1000 - 30*60).toString())}&end=${parseInt((new Date().getTime()/1000).toString())}`)
      .then(res => {
        if (res.status === 200) {
          const data = res.data.data;
          let newData:IDataInfo={
            host_name:data[0].host_name,
            cache_total: data[0].cache_total,
            connections: data[0].connections,
            ip: data[0].ip,
            request_total: data[0].request_total,
            shared_memory: data[0].shared_memory,
            status: data[0].status,
            version: data[0].version,
          };
          data.map((item: any,index:number)=>{
            if(index > 0){
              newData.host_name += " "+item.host_name;
              newData.ip += " "+item.ip;
              newData.status += " "+item.status;
              newData.version += " "+item.version;
              for(const cacheDetail in item.cache_total){
                newData.cache_total[cacheDetail] = (parseInt(newData.cache_total[cacheDetail])+parseInt(item.cache_total[cacheDetail])).toString();
              }
              for(const requestDetail in item.request_total){
                newData.request_total[requestDetail] = (parseInt(newData.request_total[requestDetail])+parseInt(item.request_total[requestDetail])).toString();
              }
  
              const itemConnections = item.connections,newDataConnections= newData.connections;
              itemConnections.map((connectionsDetail:any,index:number)=>{
                newDataConnections.map((newConnect,count) =>{
                  if(connectionsDetail.time === newConnect.time){
                    for(const timerDetail in newConnect){
                      if(timerDetail !== "time"){
                        newData.connections[count][timerDetail] = (parseInt(newConnect[timerDetail])+parseInt(connectionsDetail[timerDetail])).toString();
                      }
                    }
                  }
                  return ""
                })
                return "";
              });
              item.shared_memory.map((sharedMemoryDetail:any,index:number)=>{
                newData.shared_memory.map((newShared,count) =>{
                  if(sharedMemoryDetail.time === newShared.time){
                    for(const timerDetail in newShared){
                      if(timerDetail !== "time"){
                        newData.shared_memory[count][timerDetail] = (parseInt(newShared[timerDetail])+parseInt(sharedMemoryDetail[timerDetail])).toString();
                      }
                    }
                  }
                  return newData
                })
                return "";
              })
            }
            return '';
          })
          setDataInfo([newData]);
        } else {
          enqueueSnackbar(res.data.message, { 
            variant: 'error',
            autoHideDuration: 3000,
        });
        }
      }).catch(err => {
        enqueueSnackbar(err.response.data.message||"Error", { 
          variant: 'error',
          autoHideDuration: 3000,
      });
      })
  })
  return (
    <Layout>
      <div className={classes.root}>
        <div className={classes.block}>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Host Name</p>
            {dataInfo && dataInfo[0].host_name.split(" ").map((item,index)=>{
              return <p key={item+index} className={classes.hostDetail}>{item}</p>
              })
            }
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>IP</p>
            {dataInfo && dataInfo[0].ip.split(" ").map((item,index)=>{
              return <p key={item+index} className={classes.hostDetail}>{item}</p>
              })
            }
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Status</p>
            {dataInfo && dataInfo[0].status.split(" ").map((item,index)=>{
              return <p key={item+index} className={classes.hostDetail}>{item}</p>
              })
            }
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Version</p>
            {dataInfo && dataInfo[0].version.split(" ").map((item,index)=>{
              return <p key={item+index} className={classes.hostDetail}>{item}</p>
              })
            }
          </div>
        </div>
        <div className={classes.block}>
          <div className={classes.errorCount}>
            <p className={classes.countTitle}>In Bytes Count</p>
            <p className={classes.count}>{dataInfo?dataInfo[0].request_total.in_bytes:"0"}</p>
          </div>
          <div className={classes.errorCount}>
            <p className={classes.countTitle}>Out Bytes Count</p>
            <p className={classes.count}>{dataInfo?dataInfo[0].request_total.out_bytes:"0"}</p>
          </div>
          <div className={classes.requestCount}>
            <p className={classes.countTitle}>Request Count</p>
            <div className={classes.requestTotalContent}>
              {dataInfo && <RequestTotalCharts requestCount={dataInfo[0].request_total}/>}
            </div>
          </div>
          <div className={classes.cacheTotal}>
            <p className={classes.countTitle}>Cache Total Count</p>
            {dataInfo && <CacheTotalCharts cacheData = {dataInfo[0].cache_total}/>}
          </div>
          
        </div>
        <div className={classes.block}>
          <div className={classes.statusCodeCount}>
            <p className={classes.countTitle}>Connections Count</p>
            {dataInfo && <StatusCodeCount  connectionsData = {dataInfo[0].connections}/>}
          </div>
        </div>
        
        <div className={classes.block}>
          <div className={classes.statusCodeCount}>
            <p className={classes.countTitle}>Shared Memory Count</p>
            {dataInfo && <SharedMemoryChart  connectionsData = {dataInfo[0].shared_memory}/>}
          </div>
        </div>
      </div>
    </Layout>
  )
}
const Home = withStyles(styles)(HomeComponent);
export { Home };