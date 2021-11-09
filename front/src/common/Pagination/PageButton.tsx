import * as React from 'react';
import { memo } from 'react';

import { withStyles, StyledComponentProps } from '@material-ui/core/styles';
import classNames from 'classnames';

const styles = () => ({
  root: {
    display: 'flex',
    alignItems: 'center',
    justifyContent:'center',
    width:'360px'
  },
  pageList:{
    display: 'flex',
    alignItems: 'center',
    '& span':{
      backgroundColor: 'rgba(52,152,219,.1)',
      display:'inline-block',
      padding:'5px',
      color:'#099639',
      borderRadius:'4px',
      cursor: 'pointer',
      marginRight:'6px',
      minWidth:'38px',
      textAlign:'center',
    },
    '& .textContent':{
      color:'#8c98a4',
    }
  },
  first:{
    backgroundColor: 'rgba(52,152,219,.1)',
    display:'inline-block',
    padding:'5px',
    color:'#099639',
    borderRadius:'4px',
    cursor: 'pointer',
    marginRight:'6px',
  },
  textContent:{
    color:'#8c98a4!important',
  },
  disableded:{
    cursor: 'default',
    color:'#8c98a4!important',
  }
});

export interface IPaginationProps extends StyledComponentProps {
  page: number;
  count: number;
  perPage:number;
  onPageChange?: (page: number) => void;
}

export const scrollToTop = () => {
  window.scrollTo(0, 0);
};

const PageButtonComponent = ({
  page = 0,
  count,
  classes = {},
  onPageChange,
  perPage
}: IPaginationProps) => {
  const allPages = Math.ceil(count/perPage);
  const firstClick=()=>{
    if(onPageChange){
      if(page > 0){
        onPageChange(0);
        scrollToTop();
      }
    }
  }
  const previousClick=()=>{
    if(onPageChange){
      if(page >= 1){
        onPageChange(page-1);
        scrollToTop();
      }
      
    }
  }
  const nextClick=()=>{
    if(onPageChange){
      if(page < allPages-1){
        onPageChange(page+1);
        scrollToTop();
      }
    }
  }

  const lastClick=()=>{
    if(onPageChange){
      if(page < allPages-1){
        onPageChange(allPages-1);
        scrollToTop();
      }
    }
  }

  return (
    <div className={classes.root}>
      <div className={classes.pageList}>
         <span onClick={firstClick}  className={classNames(classes.first, page === 0 && classes.disableded)} >First</span>
         <span onClick={previousClick}  className={classNames(classes.previous, page === 0 && classes.disableded)}> &lt; </span>
         <span className={classes.textContent}>Page {page+1} of {allPages} </span>
         <span onClick={nextClick} className={classNames(classes.next, page === (allPages-1) && classes.disableded)}> &gt; </span>
         <span onClick={lastClick}  className={classNames(classes.last, page === (allPages-1) && classes.disableded)}>Last</span>
      </div>
    </div>
  );
};

export const PageButton = withStyles(styles)(memo(PageButtonComponent));
