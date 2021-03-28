import 'package:flutter/material.dart';
import 'package:flutter_datetime_picker/flutter_datetime_picker.dart';
import 'package:projeto_igreja/src/app/components/default_button_component.dart';
import 'package:projeto_igreja/src/app/components/form_error_component.dart';
import 'package:projeto_igreja/src/app/models/report.dart';
import 'package:projeto_igreja/src/app/provider/reports_provider.dart';
import 'package:projeto_igreja/src/app/provider/users_provider.dart';
import 'package:projeto_igreja/src/app/views/cellular_view/cellular_view.dart';
import 'package:projeto_igreja/src/app/views/sign_In/components/custom_svg_icon.dart';
import 'package:provider/provider.dart';

import '../../../constants.dart';
import '../../../size_config.dart';

class SendReportForm extends StatefulWidget {
  @override
  _SendReportFormState createState() => _SendReportFormState();
}

class _SendReportFormState extends State<SendReportForm> {
  final _formKey = GlobalKey<FormState>();
  final List<String> errors = [];
  final Map<String, Object> _formData = {};

  List<String> nomes = List();
  LinearGradient color = kPrimaryGradientColor;
  String place;
  String theme;
  String food;
  String date;
  bool _isDefaultPlace;
  bool _thereIsFood;
  List<String> participants = [];

  void setDefaultPlace(bool value) {
    setState(() {
      _isDefaultPlace = value;
    });
  }

  void setFoodStatus(bool value) {
    setState(() {
      _thereIsFood = value;
    });
  }

  void addError({String error}) {
    if (!errors.contains(error)) {
      setState(() {
        errors.add(error);
      });
    }
  }

  void removeError({String error}) {
    if (errors.contains(error)) {
      setState(() {
        errors.remove(error);
      });
    }
  }

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    _isDefaultPlace = true;
    _thereIsFood = true;
    _formData['date'] =
        '${DateTime.now().day} / ${DateTime.now().month} / ${DateTime.now().year}';
  }

  @override
  Widget build(BuildContext context) {
    final UsersProvider user = Provider.of(context);

    nomes = (user.all).map((item) => item.name.toString()).toList();

    return Form(
      key: _formKey,
      child: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            buildDate(context),
            SizedBox(height: getProportionateScreenHeight(10)),
            Divider(
              color: Colors.black,
            ),
            buildPlace(),
            buildFood(),
            Divider(
              color: Colors.black,
            ),
            SizedBox(height: getProportionateScreenHeight(10)),
            buildThemeFormField(),
            SizedBox(height: getProportionateScreenHeight(10)),
            Divider(
              color: Colors.black,
            ),
            Text(
              'Participantes',
              style: TextStyle(
                color: Colors.black87,
                fontSize: 25,
                fontWeight: FontWeight.w700,
              ),
            ),

            SizedBox(height: getProportionateScreenHeight(10)),

            // Presentes
            buildParticipants(user),
            SizedBox(height: getProportionateScreenHeight(10)),
            FormError(errors: errors),
            SizedBox(height: getProportionateScreenHeight(40)),
            DefaultButton(
              text: 'Salvar',
              press: () {
                if (_formKey.currentState.validate()) {
                  _formKey.currentState.save();

                  Provider.of<ReportsProvider>(context, listen: false).put(
                    Report(
                      id: _formData['id'],
                      date: _formData['date'],
                      food: _formData['food'],
                      participants: _formData['participants'],
                      place: _formData['place'],
                      theme: _formData['theme'],
                    ),
                  );
                  Navigator.pushReplacementNamed(
                      context, CellularView.routeName);
                }
              },
            ),
          ],
        ),
      ),
    );
  }

  Container buildParticipants(UsersProvider user) {
    return Container(
      height: 200,
      child: ListView.builder(
        itemCount: user.count,
        itemBuilder: (BuildContext context, int index) {
          return Stack(
            children: [
              Padding(
                padding: const EdgeInsets.symmetric(
                  vertical: 10,
                  horizontal: 10,
                ),
                child: InkWell(
                  child: Container(
                    width: double.infinity,
                    decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(24),
                        gradient: participants.contains(nomes[index])
                            ? LinearGradient(
                                begin: Alignment.topRight,
                                end: Alignment.topLeft,
                                colors: [Colors.blue, Colors.blue[800]])
                            : kPrimaryGradientColor,
                        boxShadow: [
                          BoxShadow(
                              color: kPrimaryColor.withOpacity(0.4),
                              blurRadius: 8,
                              spreadRadius: 2,
                              offset: Offset(4, 4))
                        ]),
                    child: Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 10, horizontal: 20),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        mainAxisAlignment: MainAxisAlignment.start,
                        children: [
                          Text(
                            nomes[index],
                            style: TextStyle(
                                color: Colors.white,
                                fontFamily: 'Circular',
                                fontWeight: FontWeight.w700,
                                fontSize: 25),
                          ),
                        ],
                      ),
                    ),
                  ),
                  onTap: () {
                    participants.contains(nomes[index])
                        ? setState(() {
                            participants.remove(nomes[index]);
                          })
                        : setState(() {
                            participants.add(nomes[index]);
                          });
                  },
                ),
              ),
            ],
          );
        },
      ),
    );
  }

  ElevatedButton buildDate(BuildContext context) {
    return ElevatedButton(
      style: ElevatedButton.styleFrom(
        primary: Colors.white,
        side: BorderSide(color: kPrimaryColor),
        elevation: 0.0,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(28),
        ),
      ),
      onPressed: () {
        DatePicker.showDatePicker(context,
            theme: DatePickerTheme(
              containerHeight: 210.0,
            ),
            showTitleActions: true,
            minTime: DateTime(1900, 1, 1),
            maxTime: DateTime(2050, 12, 31), onConfirm: (date) {
          _formData['date'] = '${date.day} / ${date.month} / ${date.year}';
          setState(() {});
        }, currentTime: DateTime.now(), locale: LocaleType.en);
      },
      child: Container(
        alignment: Alignment.center,
        height: 50.0,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: <Widget>[
            Row(
              children: <Widget>[
                Container(
                  child: Row(
                    children: <Widget>[
                      Icon(
                        Icons.date_range,
                        size: 18.0,
                        color: kPrimaryColor,
                      ),
                      Text(
                        "${_formData['date']}",
                        style:
                            TextStyle(color: kSecondaryColor, fontSize: 18.0),
                      ),
                    ],
                  ),
                )
              ],
            ),
            Text(
              "Data",
              style: TextStyle(color: kSecondaryColor, fontSize: 18.0),
            ),
          ],
        ),
      ),
    );
  }

  Column buildPlace() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Local:',
          style: TextStyle(
            color: Colors.black87,
            fontSize: 25,
            fontWeight: FontWeight.w700,
          ),
        ),
        Padding(
          padding: const EdgeInsets.all(0),
          child: SizedBox(
            height: 40,
            child: RadioListTile(
              activeColor: kPrimaryColor,
              value: true,
              title: Text('Local padrão de célula'),
              groupValue: _isDefaultPlace,
              onChanged: (bool value) {
                setDefaultPlace(value);
                _formData['place'] = 'Local padrão da célula';
              },
            ),
          ),
        ),
        RadioListTile(
          activeColor: kPrimaryColor,
          value: false,
          title: Text('Outro local'),
          groupValue: _isDefaultPlace,
          onChanged: (bool value) {
            setDefaultPlace(value);
            print(_isDefaultPlace);
          },
        ),
        _isDefaultPlace
            ? SizedBox(height: 10)
            : Padding(
                padding: const EdgeInsets.only(bottom: 30),
                child: TextFormField(
                  onSaved: (newValue) => _formData['place'] = newValue,
                  onChanged: (value) {
                    if (value.isNotEmpty) {
                      removeError(error: kNamelNullError);
                    }
                  },
                  validator: (value) {
                    if (value.isEmpty) {
                      addError(error: kNamelNullError);
                      return "";
                    }
                    return null;
                  },
                  decoration: InputDecoration(
                    labelText: 'Local',
                    hintText: 'Digite o local da Célula',
                    floatingLabelBehavior: FloatingLabelBehavior.always,
                    suffixIcon: CustomSuffixIcon(
                      svgIcon: 'assets/icons/hexagon.svg',
                    ),
                  ),
                ),
              )
      ],
    );
  }

  Column buildFood() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Houve lanche?',
          style: TextStyle(
            color: Colors.black87,
            fontSize: 25,
            fontWeight: FontWeight.w700,
          ),
        ),
        Padding(
          padding: const EdgeInsets.all(0),
          child: SizedBox(
            height: 40,
            child: RadioListTile(
              activeColor: kPrimaryColor,
              value: true,
              title: Text('Não'),
              groupValue: _thereIsFood,
              onChanged: (bool value) {
                setFoodStatus(value);
                _formData['food'] = 'Não houve lanche';
              },
            ),
          ),
        ),
        RadioListTile(
          activeColor: kPrimaryColor,
          value: false,
          title: Text('Sim'),
          groupValue: _thereIsFood,
          onChanged: (bool value) {
            setFoodStatus(value);
          },
        ),
        _thereIsFood
            ? SizedBox(height: 10)
            : Padding(
                padding: const EdgeInsets.only(bottom: 30),
                child: TextFormField(
                  onSaved: (newValue) => _formData['food'] = newValue,
                  onChanged: (value) {
                    if (value.isNotEmpty) {
                      removeError(error: kNamelNullError);
                    }
                  },
                  validator: (value) {
                    if (value.isEmpty) {
                      addError(error: kNamelNullError);
                      return "";
                    }
                    return null;
                  },
                  decoration: InputDecoration(
                    labelText: 'Comida',
                    hintText: 'Descreva o lanche',
                    floatingLabelBehavior: FloatingLabelBehavior.always,
                    suffixIcon: CustomSuffixIcon(
                      svgIcon: 'assets/icons/hexagon.svg',
                    ),
                  ),
                ),
              )
      ],
    );
  }

  TextFormField buildThemeFormField() {
    return TextFormField(
      onSaved: (newValue) => _formData['theme'] = newValue,
      onChanged: (value) {
        if (value.isNotEmpty) {
          removeError(error: kNamelNullError);
        }
      },
      validator: (value) {
        if (value.isEmpty) {
          addError(error: kNamelNullError);
          return "";
        }
        return null;
      },
      decoration: InputDecoration(
        labelText: 'Tema',
        hintText: 'Tema da Mensagem',
        floatingLabelBehavior: FloatingLabelBehavior.always,
        suffixIcon: CustomSuffixIcon(
          svgIcon: 'assets/icons/hexagon.svg',
        ),
      ),
    );
  }
}
