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
        { value: parseInt(requestCount._1xx||"0"), name: '_1xx' },
        { value: parseInt(requestCount._2xx||"0"), name: '_2xx' },
        { value: parseInt(requestCount._3xx||"0"), name: '_3xx' },
        { value: parseInt(requestCount._4xx||"0"), name: '_4xx' },
        { value: parseInt(requestCount._5xx||"0"), name: '_5xx' }
    ]
    const requestTotalOption: ECOption = {
        legend: {
            orient: 'left',
            left: 'right',
            textStyle:{
                color:"#fff"
            }
        },
        tooltip: {
            trigger: 'item'
        },
        series: [
            {
              name: 'Request Count',
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
                },
                label: {
                    show: true,
                    fontSize: '24',
                    fontWeight: 'bold'
                }
              },
            }
        ],
        backgroundColor:'#282c34',
    };
    return (
        <>
            <ReactEcharts option={requestTotalOption}></ReactEcharts>
        </>
    );
};

const RequestTotalCharts =   RequestTotalChartsComponent

export { RequestTotalCharts };


