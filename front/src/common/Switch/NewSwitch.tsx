import {
  StyledComponentProps,
  Switch,
  SwitchProps,
  withStyles
} from "@material-ui/core";
import React from "react";

const NewSwitch = withStyles(theme => ({
  root: {
    width: 44,
    height: 26,
    padding: 0,
    borderRadius:'14px',
  },
  switchBase: {
    padding: 1,
    "&$checked": {
      transform: "translateX(19px)",
      color: theme.palette.common.white,
      "& + $track": {
        backgroundColor: "#36C98E",
        opacity: 1,
        border: `1px solid #36C98E`
      }
    },
    "&$checked  $track": {
      backgroundColor: "#36C98E"
    },
    "&$focusVisible $thumb": {
      color: "#36C98E",
      border: "6px solid #fff"
    }
  },
  thumb: {
    margin: 2,
    width: 20,
    height: 20,
    borderRadius: '50%',
    color: theme.palette.grey[300]
  },
  track: {
    borderRadius: 4,
    border: `1px solid ${theme.palette.grey[300]}`,
    backgroundColor: theme.palette.common.white,
    opacity: 1,
    transition: theme.transitions.create(["background-color", "border"])
  },
  checked: {},
  focusVisible: {}
}))(({ classes = {}, ...props }: SwitchProps & StyledComponentProps) => {
  return (
    <Switch
      focusVisibleClassName={classes.focusVisible}
      disableRipple
      classes={{
        root: classes.root,
        switchBase: classes.switchBase,
        thumb: classes.thumb,
        track: classes.track,
        checked: classes.checked
      }}
      {...props}
    />
  );
});

export { NewSwitch };
