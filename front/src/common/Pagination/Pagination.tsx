import * as React from 'react';
import { memo } from 'react';
import { PageButton } from './PageButton'
import { withStyles, StyledComponentProps } from '@material-ui/core/styles';

const styles = () => ({
  root: {
    display: 'flex',
    alignItems: 'center',
    justifyContent:'space-between',
    paddingLeft:'24px',
  },

  changeList:{
    display: 'flex',
    alignItems: 'center',
    color:'#77838f!important',
  },
  changeSelect:{
    height:30,
    lineHeight:30,
    margin:'0 6px',
    borderRadius:'6px',
  }
});

export interface IPaginationProps extends StyledComponentProps {
  page: number;
  count: number;
  perPage: number;
  onPageChange?: (page: number) => void;
  setPerPages?: (page: number) => void;
}

const PaginationComponent = ({
  page = 0,
  count,
  perPage,
  classes = {},
  onPageChange,
  setPerPages
}: IPaginationProps) => {

  return (
    <div className={classes.root}>
        <PageButton page={page} count={count} perPage={perPage} onPageChange={onPageChange}/>
    </div>
  );
};

export const Pagination = withStyles(styles)(memo(PaginationComponent));
