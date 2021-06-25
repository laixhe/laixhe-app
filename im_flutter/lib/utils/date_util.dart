class DateUtil {

  // 时间戳
  static int timeUnix(){
    return DateTime.now().millisecondsSinceEpoch ~/ 1000;
  }

  static String? getNowDateStr() {
    return '2021-06-23';
  }

}
