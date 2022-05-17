"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CtypeFilterConds = exports.CtypeConvertables = exports.RefHardMulti = exports.RefSoftText = exports.RefHardText = exports.RefSoftPriId = exports.RefHardPriId = exports.KeyModSig = exports.KeyVersion = exports.KeyPrimary = exports.CtypeColor = exports.CtypeRangeNumber = exports.CtypeJSON = exports.CtypeEmail = exports.CtypeMultiUser = exports.CtypeSingleUser = exports.CtypeLongText = exports.CtypeMultSelect = exports.CtypeDateTime = exports.CtypeLocation = exports.CtypeNumber = exports.CtypeCurrency = exports.CtypeCheckBox = exports.CtypeMultiFile = exports.CtypeFile = exports.CtypeRFormula = exports.CtypeSelect = exports.CtypePhone = exports.CtypeShortText = void 0;
exports.CtypeShortText = "shorttext";
exports.CtypePhone = "phonenumber";
exports.CtypeSelect = "select";
exports.CtypeRFormula = "rowformula";
exports.CtypeFile = "file";
exports.CtypeMultiFile = "multifile";
exports.CtypeCheckBox = "checkbox";
exports.CtypeCurrency = "currency";
exports.CtypeNumber = "number";
exports.CtypeLocation = "location";
exports.CtypeDateTime = "datetime";
exports.CtypeMultSelect = "multiselect";
exports.CtypeLongText = "longtext";
exports.CtypeSingleUser = "singleuser";
exports.CtypeMultiUser = "multiuser";
exports.CtypeEmail = "email";
exports.CtypeJSON = "json";
exports.CtypeRangeNumber = "rangenumber";
exports.CtypeColor = "color";
// meta keys
exports.KeyPrimary = "__id";
exports.KeyVersion = "__version";
exports.KeyModSig = "__mod_sig";
exports.RefHardPriId = "hard_pri";
exports.RefSoftPriId = "soft_pri";
exports.RefHardText = "hard_text";
exports.RefSoftText = "soft_text";
exports.RefHardMulti = "hard_multi";
exports.CtypeConvertables = {
    [exports.CtypeShortText]: [exports.CtypeLongText],
    [exports.CtypePhone]: [exports.CtypeShortText],
    [exports.CtypeSelect]: [exports.CtypeShortText, exports.CtypeMultSelect],
    [exports.CtypeRFormula]: [exports.CtypeShortText],
    [exports.CtypeFile]: [exports.CtypeShortText, exports.CtypeMultiFile],
    [exports.CtypeMultiFile]: [exports.CtypeShortText],
    [exports.CtypeCheckBox]: [],
    [exports.CtypeCurrency]: [exports.CtypeNumber],
    [exports.CtypeNumber]: [exports.CtypeCurrency],
    [exports.CtypeLocation]: [],
    [exports.CtypeDateTime]: [],
    [exports.CtypeMultSelect]: [exports.CtypeShortText],
    [exports.CtypeLongText]: [exports.CtypeShortText],
    [exports.CtypeSingleUser]: [exports.CtypeShortText, exports.CtypeMultiUser],
    [exports.CtypeMultiUser]: [exports.CtypeShortText],
    [exports.CtypeEmail]: [exports.CtypeShortText],
    [exports.CtypeJSON]: [exports.CtypeShortText],
    [exports.CtypeRangeNumber]: [],
    [exports.CtypeColor]: [exports.CtypeShortText],
};
exports.CtypeFilterConds = {
    [exports.CtypeShortText]: ["equal", "not_equal", "in", "not_in"],
    [exports.CtypePhone]: ["equal", "not_equal", "in", "not_in"],
    [exports.CtypeSelect]: ["equal", "not_equal", "in", "not_in"],
    [exports.CtypeRFormula]: [],
    [exports.CtypeFile]: ["equal", "not_equal"],
    [exports.CtypeMultiFile]: [],
    [exports.CtypeCheckBox]: ["equal", "not_equal"],
    [exports.CtypeCurrency]: [
        "equal",
        "less_than",
        "not_equal",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ],
    [exports.CtypeNumber]: [
        "equal",
        "not_equal",
        "less_than",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ],
    [exports.CtypeLocation]: ["equal", "not_equal"],
    [exports.CtypeDateTime]: [
        "equal",
        "not_equal",
        "in",
        "not_in",
        "less_than",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ],
};
