import React from 'react';
import * as echarts from 'echarts/core';
import {
  BarChart,
  BarSeriesOption,
  LineSeriesOption
} from 'echarts/charts';
import {
  TitleComponent,
  TitleComponentOption,
  GridComponent,
  GridComponentOption,
  DatasetComponent,
  DatasetComponentOption,
  TransformComponent
} from 'echarts/components';
import { LabelLayout, UniversalTransition } from 'echarts/features';
import { CanvasRenderer } from 'echarts/renderers';
import ReactEcharts from "echarts-for-react";
import {StyledComponentProps} from "@material-ui/core";
import { IConnections } from "./home"
import withStyles from "@material-ui/core/styles/withStyles";
import withTheme from "@material-ui/core/styles/withTheme";
import { StyleRules } from "@material-ui/core/styles";

type ECOption = echarts.ComposeOption<
  | BarSeriesOption
  | LineSeriesOption
  | TitleComponentOption
  | GridComponentOption
  | DatasetComponentOption
>;

echarts.use([
  TitleComponent,
  GridComponent,
  DatasetComponent,
  TransformComponent,
  BarChart,
  LabelLayout,
  UniversalTransition,
  CanvasRenderer
]);

const style = (): StyleRules => ({
  chartContent:{
      height:'252px'
  }
});

interface IAppsTableRowProps extends StyledComponentProps {
  connectionsData:IConnections[];
}

const CacheTotalChartsComponent = ({
classes = {},
connectionsData,
}: IAppsTableRowProps) => {
  let legendData = [],xAxisData: string[] = [], seriesData: LineSeriesOption[] = [];
  for(const i in connectionsData[0]){
    if(i !== 'time'){
      legendData.push(i)
    }
  }
  // eslint-disable-next-line array-callback-return
  connectionsData.map((item)=>{
    xAxisData.push(item.time);
  });
  // eslint-disable-next-line array-callback-return
  legendData.map((legend)=>{
    let detail:LineSeriesOption = {
      name: legend,
      type: 'line',
      areaStyle: {
        opacity:0.3
      },
      emphasis: {
        focus: 'series'
      },
      lineStyle:{
        width:1,
      },
      symbol:'none'
    }
    let detailData: number[] = []
    connectionsData.map((item)=>{
      detailData.push(parseInt(item[legend]));
      return "";
    });
    detail['data']= detailData;
    seriesData.push(detail);
  })

  const statusCodeCountOption: ECOption = {
    legend: {
        data: legendData,
        orient: 'vertical',
        right:20,
        y:"center",
        textStyle:{
            color:"#fff"
        }
    },
    toolbox: {
        feature: {
          saveAsImage: {}
        }
    },
    xAxis: [
        {
          type: 'category',
          boundaryGap: false,
          data: xAxisData
        }
    ],
    yAxis: [
        {
          type: 'value'
        }
    ],
    series: seriesData,
    backgroundColor:'#282c34',
    grid:{
      bottom:30,
      top:20,
      right:200,
      left:80,
      height:200
    },
    tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross',
          label: {
            backgroundColor: '#6a7985'
          }
        }
    },
  };
  return (
      <div className={classes.chartContent}>
          <ReactEcharts option={statusCodeCountOption}></ReactEcharts>
      </div>
  );
};

const StatusCodeCount = withStyles(style)(
  withTheme(CacheTotalChartsComponent)
);
export { StatusCodeCount };