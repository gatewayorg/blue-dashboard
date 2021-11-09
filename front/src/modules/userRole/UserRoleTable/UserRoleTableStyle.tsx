import { StyleRules } from "@material-ui/core/styles";

const UserInfoTableStyle = (): StyleRules => ({
  status: {
    display: "flex",
    alignItems: "center"
  },

  expansionCell: {
    fontSize: 14,
    lineHeight: 1.86,
    color: "#a6adb4",
    margin: "0 4px",
    width: "50%",
    overflow: "hidden",
    wordBreak: "break-all",
    "&:first-child": { marginLeft: 0 },
    "&:last-child": { marginRight: 0 }
  },
  expansionCellTitle: {
    fontWeight: 600
  },
  appIcon: {
    width: 30,
    height: 30
  },
  buttons: {
    marginRight: 10,
    backgroundColor:'#FF4D4F',
    fontWeight: 500,
    lineHeight:'1.75',
    '&&:hover': {
      backgroundColor: '#FF4D4F',
      color:'#fff',
      opacity:'0.85'
    }
  },
  updates:{
    marginRight: 10,
    backgroundColor:'#099639',
    fontWeight: 500,
    lineHeight:'1.75',
    '&&:hover': {
      backgroundColor: '#099639',
      color:'#fff',
      opacity:'0.85'
    }
  }
});

export { UserInfoTableStyle };
