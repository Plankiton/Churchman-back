import 'package:flutter/material.dart';
import 'package:flutter_datetime_picker/flutter_datetime_picker.dart';
import 'package:projeto_igreja/src/app/components/default_button_component.dart';
import 'package:projeto_igreja/src/app/components/form_error_component.dart';
import 'package:projeto_igreja/src/app/components/theme_data.dart';
import 'package:projeto_igreja/src/app/models/user.dart';
import 'package:projeto_igreja/src/app/provider/users_provider.dart';
import 'package:projeto_igreja/src/app/views/cellular_view/cellular_view.dart';
import 'package:projeto_igreja/src/app/views/sign_In/components/custom_svg_icon.dart';
import 'package:provider/provider.dart';

import '../../../constants.dart';
import '../../../size_config.dart';

class GanharForm extends StatefulWidget {
  @override
  _GanharFormState createState() => _GanharFormState();
}

class _GanharFormState extends State<GanharForm> {
  final _formKey = GlobalKey<FormState>();
  final List<String> errors = [];
  final Map<String, Object> _formData = {};

  String firstName;
  String lastName;
  String phoneNumber;
  String address;
  String sexo;
  String civilState;
  String born = '1 / 1 / 2000 ';

  List listSexo = [
    'Masculino',
    'Feminino',
  ];

  List listCivil = [
    'Casado(a)',
    'Solteiro(a)',
  ];

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
    _formData['born'] =
        '${DateTime.now().day} / ${DateTime.now().month} / ${DateTime.now().year}';
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: SingleChildScrollView(
        child: Column(
          children: [
            buildFirstNameFormField(),
            SizedBox(height: getProportionateScreenHeight(30)),
            buildPhoneNumberFormField(),
            SizedBox(height: getProportionateScreenHeight(30)),
            // buildAddressFormField(),
            // SizedBox(height: getProportionateScreenHeight(30)),
            buildCivilSexoList(),
            SizedBox(height: getProportionateScreenHeight(30)),
            buildBorn(context),
            SizedBox(height: getProportionateScreenHeight(10)),
            FormError(errors: errors),
            SizedBox(height: getProportionateScreenHeight(40)),
            DefaultButton(
              text: 'Salvar',
              press: () {
                if (_formKey.currentState.validate()) {
                  _formKey.currentState.save();

                  Provider.of<UsersProvider>(context, listen: false).put(
                    User(
                      id: _formData['id'],
                      name: _formData['name'],
                      email: _formData['email'],
                      data: _formData['data'],
                      sex: _formData['sex'],
                      phone: _formData['phone'],
                      born: _formData['born'],
                      state: _formData['state'],
                    ),
                  );
                  //Vai para Home
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

  Row buildCivilSexoList() {
    return Row(
      mainAxisSize: MainAxisSize.min,
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        //Sexo
        Expanded(
          flex: 1,
          child: Container(
            width: getProportionateScreenWidth(20),
            height: getProportionateScreenHeight(50),
            child: DropdownButtonFormField(
              onSaved: (newValue) => _formData['sex'] = newValue,
              icon: Icon(Icons.arrow_drop_down_rounded),
              iconSize: getProportionateScreenHeight(30),
              iconDisabledColor: kPrimaryColor,
              iconEnabledColor: kPrimaryColor,

              style: TextStyle(
                color: kSecondaryColor,
                fontSize: getProportionateScreenHeight(18),
              ),

              value: _formData['sex'],

              items: listSexo.map((valueItem) {
                return DropdownMenuItem(
                  value: valueItem,
                  child: Text(valueItem, overflow: TextOverflow.visible),
                );
              }).toList(), //items
              hint: Text('Sexo'),

              onChanged: (value) {
                print(value);
                if (value != 'Sexo') {
                  removeError(error: kSexNullError);
                }
              },
              validator: (value) {
                print(value);
                if (value == 'Sexo') {
                  addError(error: kSexNullError);
                  return "";
                }
                return null;
              },
              decoration: inputDecorationRow(labelText: 'Sexo'),
            ),
          ),
        ),
        //ESTADO CIVIL
        SizedBox(width: getProportionateScreenWidth(20)),
        Expanded(
          flex: 1,
          child: Container(
            width: getProportionateScreenWidth(20),
            height: getProportionateScreenHeight(50),
            child: DropdownButtonFormField(
              onSaved: (newValue) => _formData['state'] = newValue,
              icon: Icon(Icons.arrow_drop_down_rounded),
              iconSize: getProportionateScreenHeight(30),
              iconDisabledColor: kPrimaryColor,
              iconEnabledColor: kPrimaryColor,

              style: TextStyle(
                color: kSecondaryColor,
                fontSize: getProportionateScreenHeight(18),
              ),

              value: _formData['state'],

              items: listCivil.map((valueItem) {
                return DropdownMenuItem(
                  value: valueItem,
                  child: Text(valueItem, overflow: TextOverflow.visible),
                );
              }).toList(), //items

              hint: Text('Estado Civil'),

              onChanged: (value) {
                if (value != 'Estado Civil') {
                  removeError(error: kCivilNullError);
                }
              },
              validator: (value) {
                if (value == 'Estado Civil') {
                  addError(error: kCivilNullError);
                  return "";
                }
                return null;
              },

              decoration: inputDecorationRow(labelText: 'Estado Civl'),
            ),
          ),
        ),
      ],
    );
  }

  TextFormField buildFirstNameFormField() {
    return TextFormField(
      onSaved: (newValue) => _formData['name'] = newValue,
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
        labelText: 'Nome',
        hintText: 'Digite seu Nome',
        floatingLabelBehavior: FloatingLabelBehavior.always,
        suffixIcon: CustomSuffixIcon(
          svgIcon: 'assets/icons/User.svg',
        ),
      ),
    );
  }

  TextFormField buildPhoneNumberFormField() {
    return TextFormField(
      keyboardType: TextInputType.number,
      onSaved: (newValue) => _formData['phone'] = newValue,
      onChanged: (value) {
        if (value.isNotEmpty) {
          removeError(error: kPhoneNumberNullError);
        }
      },
      validator: (value) {
        if (value.isEmpty) {
          addError(error: kPhoneNumberNullError);
          return "";
        }
        return null;
      },
      decoration: InputDecoration(
        labelText: 'Telefone',
        hintText: 'Digite seu Telefone',
        floatingLabelBehavior: FloatingLabelBehavior.always,
        suffixIcon: CustomSuffixIcon(
          svgIcon: 'assets/icons/Phone.svg',
        ),
      ),
    );
  }

  ElevatedButton buildBorn(BuildContext context) {
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
          _formData['born'] = '${date.day} / ${date.month} / ${date.year}';
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
                        "${_formData['born']}",
                        style:
                            TextStyle(color: kSecondaryColor, fontSize: 18.0),
                      ),
                    ],
                  ),
                )
              ],
            ),
            Text(
              "Nascimento",
              style: TextStyle(color: kSecondaryColor, fontSize: 18.0),
            ),
          ],
        ),
      ),
    );
  }
}
