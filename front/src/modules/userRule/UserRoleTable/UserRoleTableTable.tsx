import React from "react";
import {
  StyledComponentProps,
  Theme,
  Button
} from "@material-ui/core";
import { t } from '../../../common/utils/intl';
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import withTheme from "@material-ui/core/styles/withTheme";
import withStyles from "@material-ui/core/styles/withStyles";
import { UserInfoTableStyle } from "./UserRoleTableStyle";
import { IDataInfo } from '../userRule';

interface IAppsTableRowProps extends StyledComponentProps {
  theme: Theme;
  datadetail: IDataInfo;
  setUserSelect: React.Dispatch<React.SetStateAction<IDataInfo | undefined>>;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;

}

const UserRoleTableComponent = ({
  classes = {},
  theme,
  datadetail,
  setUserSelect,
  setOpen,
}: IAppsTableRowProps) => {

  const setRule  = ()=>{
    setOpen(true);
    setUserSelect(datadetail);
  }

  return (
    <>
      <TableRow
        key={datadetail.id}
        hover
        style={{ cursor: "pointer" }}
      >
        <TableCell>{datadetail.id}</TableCell>
        <TableCell>{datadetail.method}</TableCell>
        <TableCell>{datadetail.service}</TableCell>
        <TableCell>{datadetail.detail}</TableCell>
        <TableCell>
        <Button onClick={setRule} className={classes.buttons}>
          {t("user-rule.header.set")}
        </Button>
        </TableCell>
      </TableRow>
    </>
  );
};

const UserRoleTable = withStyles(UserInfoTableStyle)(
  withTheme(UserRoleTableComponent)
);

export { UserRoleTable };
