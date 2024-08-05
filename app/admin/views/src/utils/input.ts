import { InputTypeList } from 'src/utils/define';

// executeEval 执行eval代码
export const executeEval = (evalStr: string, data: { scope: any }) => {
  if (evalStr == '') return true;
  const { scope } = data;
  console.log(scope.row)
  return eval(evalStr);
}

export const inputOperateSettingFun = (
  operate: string,
  tableConfig: any,
  btnConfig: any,
  rawData: any
) => {
  const operateValue = {} as any;
  operateValue.params = {} as any;
  operateValue.params[tableConfig.table.key] =
    rawData.row[tableConfig.table.key];
  operateValue.params['type'] = rawData.row[btnConfig.params['type']];

  //  默认输入框列表
  operateValue.inputList = [] as any;
  operateValue.inputList.push({
    label: rawData.row[btnConfig.params['name']],
    field: btnConfig.params['value'],
    type: rawData.row[btnConfig.params['type']],
    data: rawData.row[btnConfig.params['input']],
  });

  //  处理值问题
  switch (rawData.row[btnConfig.params['type']]) {
    //  多选框值初始化
    case InputTypeList.Checkbox:
    case InputTypeList.InputJson:
    case InputTypeList.InputChildren:
    case InputTypeList.Images:
      operateValue.params['value'] = JSON.parse(
        rawData.row[btnConfig.params['value']]
      );

      if (rawData.row[btnConfig.params['input']] != '') {
        operateValue.inputList[0].data = JSON.parse(
          rawData.row[btnConfig.params['input']]
        );
      }

      break;
    case InputTypeList.Select:
      operateValue.params['value'] = rawData.row[btnConfig.params['value']];
      if (rawData.row[btnConfig.params['input']] != '') {
        operateValue.inputList[0].data = JSON.parse(
          rawData.row[btnConfig.params['input']]
        );
      }
      break;
    //  其他类型值设置
    default:
      operateValue.params['value'] = rawData.row[btnConfig.params['value']];
      break;
  }

  return operateValue;
};
