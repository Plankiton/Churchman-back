import 'package:flutter/material.dart';
import 'package:projeto_igreja/src/app/components/default_open_text.dart';
import 'package:projeto_igreja/src/app/size_config.dart';
import 'package:projeto_igreja/src/app/views/send_report/components/send_report_form.dart';

class SendReportBody extends StatefulWidget {
  @override
  _SendReportBodyState createState() => _SendReportBodyState();
}

class _SendReportBodyState extends State<SendReportBody> {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Padding(
        padding:
            EdgeInsets.symmetric(horizontal: getProportionateScreenWidth(20)),
        child: SingleChildScrollView(
          child: Column(
            children: [
              SizedBox(height: SizeConfig.screenHeight * 0.02),
              DefaultOpenText(
                  title: 'Relatório de Célula',
                  subtitle: 'Complete com as informações\nreferentes à célula'),
              SizedBox(height: SizeConfig.screenHeight * 0.05),
              SendReportForm(),
            ],
          ),
        ),
      ),
    );
  }
}
