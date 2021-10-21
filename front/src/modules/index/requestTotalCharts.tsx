import React from 'react';
import * as echarts from 'echarts/core';
import {
  BarChart,
  BarSeriesOption,
  LineSeriesOption,
  PieSeriesOption
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
import { IRequestTotal } from "./home"

type ECOption = echarts.ComposeOption<
  | BarSeriesOption
  | LineSeriesOption
  | TitleComponentOption
  | GridComponentOption
  | DatasetComponentOption
  | PieSeriesOption
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

interface IAppsTableRowProps extends StyledComponentProps {
    requestCount:IRequestTotal;
}

const RequestTotalChartsComponent = ({
  classes = {},
  requestCount,
}: IAppsTableRowProps) => {
    const chartData  = [
        { value: parseInt(requestCount._1xx), name: '_1xx' },
        { value: parseInt(requestCount._2xx), name: '_2xx' },
        { value: parseInt(requestCount._3xx), name: '_3xx' },
        { value: parseInt(requestCount._4xx), name: '_4xx' },
        { value: parseInt(requestCount._5xx), name: '_5xx' }
    ]
    const requestTotalOption: ECOption = {
        legend: {
            orient: 'left',
            left: 'right',
            textStyle:{
                color:"#fff"
            }
        },
        series: [
            {
              name: 'Access From',
              type: 'pie',
              radius: '50%',
              label:{
                  show:false
              },
              center:["35%","35%"],
              data: chartData,
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              }
            }
        ],
        backgroundColor:'#282c34',
        label:{
            normal:{
                show:false,
                position: 'inside',
                formatter:"{b}:{d}%"
            }
        },
    
    };
    return (
        <>
            <ReactEcharts option={requestTotalOption}></ReactEcharts>
        </>
    );
};

const RequestTotalCharts =   RequestTotalChartsComponent

export { RequestTotalCharts };


