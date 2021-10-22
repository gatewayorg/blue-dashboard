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
import { client } from '../../api/service';
import { useInitEffect } from '../../hooks/useInitEffect';
import { CacheTotalCharts } from "./cacheTotalCharts";
import { SharedMemoryChart } from "./sharedMemoryChart"

const styles = (theme: Theme): StyleRules => ({
  root:{
    
  },
  block: {
    display:"flex",
    margin:"20px 20px",
    borderRadius:"10px",
    overflow:'hidden'
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
  },
  requestCount:{
    backgroundColor:'#282c34',
    marginLeft:'10px',
    borderRadius:'8px',
    width:"20%",
    height:300,
    textAlign:'center',
    border: '1px solid rgba(256,256,256,0.1)',
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
  },
  statusCodeCount:{
    width:"calc(100% - 20px)",
    border: '1px solid rgba(256,256,256,0.2)',
    borderRadius:"8px",

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
    overflow:"hidden"
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
  },
  hostDetail:{
    color:'#099639',
    fontSize:"22px",
    textAlign:'center',
    margin:0,
    fontWeight:500,
    lineHeight:'160px',
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
  bypass?: string;
  expired?:string;
  hit?:string;
  miss?:string;
  revalidated?:string;
  scarce?:string;
  stale?:string;
  updating?:string;
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

  useInitEffect(() => {
    client.get(`/index?start=${parseInt((new Date().getTime()/1000 - 24*60*60).toString())}&end=${parseInt((new Date().getTime()/1000).toString())}`)
      .then(res => {
        if (res.status === 200) {
          const data = res.data.data;
          setDataInfo(data);
          console.log(data);

        }
      })
  })
  return (
    <Layout>
      <div className={classes.root}>
        <div className={classes.block}>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Host Name</p>
            <p className={classes.hostDetail}>{dataInfo?dataInfo[0].host_name:""}</p>
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>IP</p>
            <p className={classes.hostDetail}>{dataInfo?dataInfo[0].ip:""}</p>
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Status</p>
            <p className={classes.hostDetail}>{dataInfo?dataInfo[0].status:""}</p>
          </div>
          <div className={classes.hostInfo}>
            <p className={classes.countTitle}>Version</p>
            <p className={classes.hostDetail}>{dataInfo?dataInfo[0].version:""}</p>
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