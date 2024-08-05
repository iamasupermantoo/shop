import { exportFile } from 'quasar';
import { WarningNotify } from 'src/utils/notify';

// 下载csv文件
export const exportCSVFile = (columns: any, data: any) => {
  const content = [columns.map((col: any) => col.label)]
    .concat(
      data.map((row: any) =>
        columns
          .map(
            (col: any) =>
              '"' +
              row[col.field] ? row[col.field] : '' +
              '"'
          )
          .join(',')
      )
    )
    .join('\r\n');

  const status = exportFile('tables.csv', content, 'text/csv');

  if (status !== true) {
    WarningNotify('Browser denied file download...');
  }
};
