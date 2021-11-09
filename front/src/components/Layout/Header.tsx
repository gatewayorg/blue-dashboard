import React from 'react';
import { t } from '../../common/utils/intl';
import { 
  StyledComponentProps,
  StyleRules, 
  Theme, 
  withStyles
} from '@material-ui/core';
import { useInitEffect } from '../../hooks/useInitEffect';
// import { ReactComponent as Logo } from '../../assets/logo.svg';
import { PATH } from '../../nav/const';
import { Link as RouterLink } from 'react-router-dom';


const styles = (theme: Theme): StyleRules => ({
  root: {
    width: '100%', 
    backgroundColor: '#000',  
    padding:'0 60px',
  },
  logo:{
    height:40,
  },
  headers:{
    width: '100%',
    height: theme.spacing(10),
    margin:'0 auto',
    display:'flex',
    alignItems:'center',
    justifyContent:'flex-start'
  },
  link:{
    display:'flex',
    height:'100%',
    alignItems:'center'
  },
  route:{
    padding:'0 30px',
    height:'100%',
    display:'block',
    lineHeight:'80px',
    textDecoration:'none',
    fontFamily: 'Noto Sans S Chinese',
    fontStyle: 'normal',
    fontWeight:500,
    fontSize:'16px',
    color:'rgba(255, 255, 255, 0.5)',
    '&&:hover': {
      color:'rgba(255, 255, 255, 0.9)',
    }
  },
  content:{
    margin:'0 30px',
    height:'100%',
    display:'block',
    lineHeight:'80px',
    textDecoration:'none',
    fontFamily: 'Noto Sans S Chinese',
    fontStyle: 'normal',
    fontWeight:500,
    fontSize:'16px',
    color:'rgba(255, 255, 255, 0.5)',
    cursor:'pointer',
    position:'relative',
    '&:hover $contentRoute': {
      display:'block',
    }
  },
  contentRoute:{
    padding:'0 30px',
    height:'44px',
    lineHeight:'44px',
    textDecoration:'none',
    fontFamily: 'Noto Sans S Chinese',
    fontStyle: 'normal',
    fontWeight:500,
    fontSize:'16px',
    color:'rgba(255, 255, 255, 0.5)',
    backgroundColor:'#000',
    marginLeft:'-30px',
    '&&:hover': {
      color:'rgba(255, 255, 255, 0.9)',
    },
    display:'none',
    whiteSpace:"nowrap",
  }
});

interface IHeaderProps
extends StyledComponentProps {

}

const HeaderComponent = ({
  classes = {},
}: IHeaderProps) => {



  useInitEffect(() => {

  })




  return (
    <div className={classes.root}>
      <div className={classes.headers}>
          <RouterLink to={PATH.DASHBOARD}>
            {/* <Logo className={classes.logo} /> */}
          </RouterLink>
          <div className={classes.link}>
            <RouterLink to={PATH.DASHBOARD} className={classes.route}>
              {t('header.home')}
            </RouterLink>
            <div className={classes.content}>
              {t('header.my')}
              <div>
                <RouterLink to={PATH.MY.My} className={classes.contentRoute}>
                  {t('header.info')}
                </RouterLink>
                <RouterLink to={PATH.MY.Rule} className={classes.contentRoute}>
                  {t('header.rule')}
                </RouterLink>
                <RouterLink to={PATH.MY.Role} className={classes.contentRoute}>
                  {t('header.role')}
                </RouterLink>
              </div>
            </div> 
          </div>
          <div>
            {/* <p>123</p> */}
          </div>
        </div>
    </div>
  )
}

const Header = withStyles(styles)(HeaderComponent);

export { Header };