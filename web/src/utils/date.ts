const unix2Date = (u: number) => {
  const temp_time = new Date(u * 1000);
  const year = temp_time.getFullYear();
  //取得日期中的月份，其中0表示1月，11表示12月
  let month = temp_time.getMonth() + 1;
  //小于10月的月份补全0 例如1月补全为01月
  month = month < 10 ? "0" + month : month;
  //返回日期月份中的天数（1到31）
  let day = temp_time.getDate();
  day = day < 10 ? "0" + day : day;
  //返回日期中的小时数（0到23）
  let hour = temp_time.getHours();
  hour = hour < 10 ? "0" + hour : hour;
  //返回日期中的分钟数（0到59）
  let minute = temp_time.getMinutes();
  minute = minute < 10 ? "0" + minute : minute;
  const result_time =
    year + "-" + month + "-" + day + " " + hour + ":" + minute;
  return result_time;
};

export default unix2Date;
