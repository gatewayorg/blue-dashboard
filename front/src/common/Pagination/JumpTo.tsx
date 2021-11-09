import * as React from 'react';
import { withStyles, StyledComponentProps } from '@material-ui/core/styles';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { t } from '../utils/intl';

const styles = () => ({
  button: {
    padding: '2px 12px',
    borderRadius: 4,
    display: 'flex',
    alignItems: 'center',
    cursor: 'pointer',
    '& .MuiInputBase-root':{
      margin:0,
      width: 70,
      textAlign: 'center',
      height:35,
      backgroundColor:'rgb(90 90 91)',
    },
    '& .MuiSelect-select.MuiSelect-select':{
      backgroundColor:'rgb(90 90 91)',
      color:'#fff',
      fontSize:'20px',
      padding:0,
    }
  },
  label: {
    fontSize: 21,
    lineHeight: 1.57,
    color: '#fff',
    marginRight: 8,
  },
  input: {
    border: 'none',
    fontSize: 14,
    padding: 0,
    lineHeight: 1.57,
    color: 'rgba(0, 0, 0, 0.65)',
    minWidth: 'auto',
  },
  page: {
    fontSize: 21,
    lineHeight: 1.57,
    color: '#fff',
    marginLeft: 8,
  },
});

export interface IJumpToProps extends StyledComponentProps {
  page: number;
  pageCount: number;
  onPageChange?: (page: number) => void;
  disabled?: boolean;
}

const JumpToComponent = ({
  page = 0,
  pageCount,
  disabled = false,
  onPageChange,
  classes = {},
}: IJumpToProps) => {
  const [open, setOpen] = React.useState(false);

  const handleSelect = () => {
    setOpen(!open);
  };

  const handleChange = (event: any) => {
    let value: number = event.target.value;
    onPageChange && !disabled && onPageChange(value);
  };

  const renderIcon = () => <div />;

  return (
    <>
      <div className={classes.button} onClick={handleSelect}>
        <div className={classes.label}>{t('pagination.jump-to')}</div>
        <Select
          disableUnderline
          value={page}
          open={open}
          onChange={handleChange}
          classes={{
            select: classes.input,
          }}
          IconComponent={renderIcon}
        >
          {[...Array(pageCount)].map((item, i) => {
            const page = i + 1;
            return (
              <MenuItem value={i} key={page}>
                {page}
              </MenuItem>
            );
          })}
        </Select>
      </div>

      <div className={classes.page}>{t('pagination.page')}</div>
    </>
  );
};

export const JumpTo = withStyles(styles)(JumpToComponent);
