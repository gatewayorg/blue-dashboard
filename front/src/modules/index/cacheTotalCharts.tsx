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
import withStyles from "@material-ui/core/styles/withStyles";
import withTheme from "@material-ui/core/styles/withTheme";
import { StyleRules } from "@material-ui/core/styles";
import { ICacheTotal } from "./home"

const style = (): StyleRules => ({
    chartContent:{
        height:'252px'
    }
  });

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
    cacheData:ICacheTotal;
}

const CacheTotalChartsComponent = ({
  classes = {},
  cacheData,
}: IAppsTableRowProps) => {
    let dataList = []
    for(const i in cacheData){
        dataList.push(i)
    }
    let seriesData:number[] = [];
    // eslint-disable-next-line array-callback-return
    dataList.map((item) => {
        seriesData.push(parseInt(cacheData[item]))
    })

    const requestTotalOption: ECOption = {
        xAxis: {
            type: 'category',
            data: dataList,
            axisLabel:{interval: 0}
        },
        yAxis: {
            type: 'value'
        },
        color:["#099639"],
        series: [
            {
              data: seriesData,
              type: 'bar'
            }
        ],
        backgroundColor:'#282c34',
        grid:{
            bottom:30,
            top:20,
            right:20,
            left:50,
            height:200
        },
    };
    return (
        <div className={classes.chartContent}>
            <ReactEcharts option={requestTotalOption}></ReactEcharts>
        </div>
    );
};

const CacheTotalCharts = withStyles(style)(
    withTheme(CacheTotalChartsComponent)
  );
export { CacheTotalCharts };


