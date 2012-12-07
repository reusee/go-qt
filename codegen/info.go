package main

var classes = make(map[string]*Class)
var types = make(map[int]*Type)

var klass *Class
var type_ *Type

func init() {
	klass = &Class{
		"QAbstractAnimation",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractAnimation"] = klass
	klass = &Class{
		"QAbstractEventDispatcher",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractEventDispatcher"] = klass
	klass = &Class{
		"QAbstractFileEngine",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractFileEngine"] = klass
	klass = &Class{
		"QAbstractFileEngine::ExtensionOption",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractFileEngine::ExtensionOption"] = klass
	klass = &Class{
		"QAbstractFileEngine::ExtensionReturn",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractFileEngine::ExtensionReturn"] = klass
	klass = &Class{
		"QAbstractFileEngine::MapExtensionOption",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractFileEngine::MapExtensionOption"] = klass
	klass = &Class{
		"QAbstractFileEngine::MapExtensionReturn",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractFileEngine::MapExtensionReturn"] = klass
	klass = &Class{
		"QAbstractFileEngine::UnMapExtensionOption",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractFileEngine::UnMapExtensionOption"] = klass
	klass = &Class{
		"QAbstractFileEngineHandler",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAbstractFileEngineHandler"] = klass
	klass = &Class{
		"QAbstractFileEngineIterator",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractFileEngineIterator"] = klass
	klass = &Class{
		"QAbstractItemModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractItemModel"] = klass
	klass = &Class{
		"QAbstractListModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractListModel"] = klass
	klass = &Class{
		"QAbstractState",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractState"] = klass
	klass = &Class{
		"QAbstractTableModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractTableModel"] = klass
	klass = &Class{
		"QAbstractTransition",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractTransition"] = klass
	klass = &Class{
		"QAnimationGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAnimationGroup"] = klass
	klass = &Class{
		"QAtomicInt",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAtomicInt"] = klass
	klass = &Class{
		"QBasicAtomicInt",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBasicAtomicInt"] = klass
	klass = &Class{
		"QBasicTimer",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBasicTimer"] = klass
	klass = &Class{
		"QBitArray",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBitArray"] = klass
	klass = &Class{
		"QBitRef",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBitRef"] = klass
	klass = &Class{
		"QBool",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBool"] = klass
	klass = &Class{
		"QBuffer",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QBuffer"] = klass
	klass = &Class{
		"QByteArray",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QByteArray"] = klass
	klass = &Class{
		"QByteArrayMatcher",
		false,
		make([]*Class, 0),
		1040,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QByteArrayMatcher"] = klass
	klass = &Class{
		"QByteRef",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QByteRef"] = klass
	klass = &Class{
		"QChar",
		false,
		make([]*Class, 0),
		2,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QChar"] = klass
	klass = &Class{
		"QCharRef",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QCharRef"] = klass
	klass = &Class{
		"QChildEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QChildEvent"] = klass
	klass = &Class{
		"QCoreApplication",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCoreApplication"] = klass
	klass = &Class{
		"QCryptographicHash",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QCryptographicHash"] = klass
	klass = &Class{
		"QDataStream",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDataStream"] = klass
	klass = &Class{
		"QDate",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QDate"] = klass
	klass = &Class{
		"QDateTime",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QDateTime"] = klass
	klass = &Class{
		"QDebug",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QDebug"] = klass
	klass = &Class{
		"QDir",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QDir"] = klass
	klass = &Class{
		"QDirIterator",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDirIterator"] = klass
	klass = &Class{
		"QDynamicPropertyChangeEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDynamicPropertyChangeEvent"] = klass
	klass = &Class{
		"QEasingCurve",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QEasingCurve"] = klass
	klass = &Class{
		"QElapsedTimer",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QElapsedTimer"] = klass
	klass = &Class{
		"QEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QEvent"] = klass
	klass = &Class{
		"QEventLoop",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QEventLoop"] = klass
	klass = &Class{
		"QEventPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QEventPrivate"] = klass
	klass = &Class{
		"QEventTransition",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QEventTransition"] = klass
	klass = &Class{
		"QFSFileEngine",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFSFileEngine"] = klass
	klass = &Class{
		"QFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QFactoryInterface"] = klass
	klass = &Class{
		"QFile",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFile"] = klass
	klass = &Class{
		"QFileInfo",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFileInfo"] = klass
	klass = &Class{
		"QFileSystemWatcher",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFileSystemWatcher"] = klass
	klass = &Class{
		"QFinalState",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFinalState"] = klass
	klass = &Class{
		"QFlag",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFlag"] = klass
	klass = &Class{
		"QFutureInterfaceBase",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QFutureInterfaceBase"] = klass
	klass = &Class{
		"QFutureWatcherBase",
		false,
		make([]*Class, 0),
		16,
		false,
		false,
		true,
		false,
		false,
	}
	classes["QFutureWatcherBase"] = klass
	klass = &Class{
		"QGenericArgument",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QGenericArgument"] = klass
	klass = &Class{
		"QGenericReturnArgument",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QGenericReturnArgument"] = klass
	klass = &Class{
		"QGlobalSpace",
		false,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		true,
		false,
	}
	classes["QGlobalSpace"] = klass
	klass = &Class{
		"QHashDummyValue",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QHashDummyValue"] = klass
	klass = &Class{
		"QHistoryState",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QHistoryState"] = klass
	klass = &Class{
		"QIODevice",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QIODevice"] = klass
	klass = &Class{
		"QIncompatibleFlag",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QIncompatibleFlag"] = klass
	klass = &Class{
		"QInternal",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QInternal"] = klass
	klass = &Class{
		"QLatin1Char",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLatin1Char"] = klass
	klass = &Class{
		"QLatin1String",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLatin1String"] = klass
	klass = &Class{
		"QLibrary",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QLibrary"] = klass
	klass = &Class{
		"QLibraryInfo",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLibraryInfo"] = klass
	klass = &Class{
		"QLine",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLine"] = klass
	klass = &Class{
		"QLineF",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLineF"] = klass
	klass = &Class{
		"QLocale",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLocale"] = klass
	klass = &Class{
		"QMargins",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMargins"] = klass
	klass = &Class{
		"QMetaClassInfo",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaClassInfo"] = klass
	klass = &Class{
		"QMetaEnum",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaEnum"] = klass
	klass = &Class{
		"QMetaMethod",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaMethod"] = klass
	klass = &Class{
		"QMetaObject",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaObject"] = klass
	klass = &Class{
		"QMetaProperty",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaProperty"] = klass
	klass = &Class{
		"QMetaType",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMetaType"] = klass
	klass = &Class{
		"QMimeData",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMimeData"] = klass
	klass = &Class{
		"QModelIndex",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QModelIndex"] = klass
	klass = &Class{
		"QMutex",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QMutex"] = klass
	klass = &Class{
		"QNoDebug",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QNoDebug"] = klass
	klass = &Class{
		"QObject",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QObject"] = klass
	klass = &Class{
		"QObjectCleanupHandler",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QObjectCleanupHandler"] = klass
	klass = &Class{
		"QObjectUserData",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QObjectUserData"] = klass
	klass = &Class{
		"QParallelAnimationGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QParallelAnimationGroup"] = klass
	klass = &Class{
		"QPauseAnimation",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPauseAnimation"] = klass
	klass = &Class{
		"QPersistentModelIndex",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPersistentModelIndex"] = klass
	klass = &Class{
		"QPluginLoader",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPluginLoader"] = klass
	klass = &Class{
		"QPoint",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPoint"] = klass
	klass = &Class{
		"QPointF",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPointF"] = klass
	klass = &Class{
		"QPostEventList",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPostEventList"] = klass
	klass = &Class{
		"QProcess",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QProcess"] = klass
	klass = &Class{
		"QProcessEnvironment",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QProcessEnvironment"] = klass
	klass = &Class{
		"QPropertyAnimation",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPropertyAnimation"] = klass
	klass = &Class{
		"QReadLocker",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QReadLocker"] = klass
	klass = &Class{
		"QReadWriteLock",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QReadWriteLock"] = klass
	klass = &Class{
		"QRect",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRect"] = klass
	klass = &Class{
		"QRectF",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRectF"] = klass
	klass = &Class{
		"QRegExp",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRegExp"] = klass
	klass = &Class{
		"QResource",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QResource"] = klass
	klass = &Class{
		"QRunnable",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QRunnable"] = klass
	klass = &Class{
		"QSemaphore",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QSemaphore"] = klass
	klass = &Class{
		"QSequentialAnimationGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSequentialAnimationGroup"] = klass
	klass = &Class{
		"QSettings",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSettings"] = klass
	klass = &Class{
		"QSharedData",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QSharedData"] = klass
	klass = &Class{
		"QSharedMemory",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSharedMemory"] = klass
	klass = &Class{
		"QSignalMapper",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSignalMapper"] = klass
	klass = &Class{
		"QSignalTransition",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSignalTransition"] = klass
	klass = &Class{
		"QSize",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QSize"] = klass
	klass = &Class{
		"QSizeF",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QSizeF"] = klass
	klass = &Class{
		"QSocketNotifier",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSocketNotifier"] = klass
	klass = &Class{
		"QState",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QState"] = klass
	klass = &Class{
		"QStateMachine",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStateMachine"] = klass
	klass = &Class{
		"QString::Null",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QString::Null"] = klass
	klass = &Class{
		"QStringMatcher",
		false,
		make([]*Class, 0),
		1048,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStringMatcher"] = klass
	klass = &Class{
		"QStringRef",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStringRef"] = klass
	klass = &Class{
		"QSysInfo",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QSysInfo"] = klass
	klass = &Class{
		"QSystemLocale",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QSystemLocale"] = klass
	klass = &Class{
		"QSystemSemaphore",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QSystemSemaphore"] = klass
	klass = &Class{
		"QTemporaryFile",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTemporaryFile"] = klass
	klass = &Class{
		"QTextBoundaryFinder",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextBoundaryFinder"] = klass
	klass = &Class{
		"QTextCodec",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextCodec"] = klass
	klass = &Class{
		"QTextCodec::ConverterState",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QTextCodec::ConverterState"] = klass
	klass = &Class{
		"QTextCodecFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTextCodecFactoryInterface"] = klass
	klass = &Class{
		"QTextCodecPlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextCodecPlugin"] = klass
	klass = &Class{
		"QTextDecoder",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QTextDecoder"] = klass
	klass = &Class{
		"QTextEncoder",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QTextEncoder"] = klass
	klass = &Class{
		"QTextStream",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextStream"] = klass
	klass = &Class{
		"QTextStreamManipulator",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextStreamManipulator"] = klass
	klass = &Class{
		"QThread",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QThread"] = klass
	klass = &Class{
		"QTime",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTime"] = klass
	klass = &Class{
		"QTimeLine",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTimeLine"] = klass
	klass = &Class{
		"QTimer",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTimer"] = klass
	klass = &Class{
		"QTimerEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTimerEvent"] = klass
	klass = &Class{
		"QTranslator",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTranslator"] = klass
	klass = &Class{
		"QUrl",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QUrl"] = klass
	klass = &Class{
		"QUrlPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QUrlPrivate"] = klass
	klass = &Class{
		"QUuid",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QUuid"] = klass
	klass = &Class{
		"QVariant",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QVariant"] = klass
	klass = &Class{
		"QVariant::Handler",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QVariant::Handler"] = klass
	klass = &Class{
		"QVariant::Private",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QVariant::Private"] = klass
	klass = &Class{
		"QVariantAnimation",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QVariantAnimation"] = klass
	klass = &Class{
		"QVariantComparisonHelper",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QVariantComparisonHelper"] = klass
	klass = &Class{
		"QWidget",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QWidget"] = klass
	klass = &Class{
		"QWriteLocker",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QWriteLocker"] = klass
	klass = &Class{
		"QXmlStreamAttribute",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamAttribute"] = klass
	klass = &Class{
		"QXmlStreamAttributes",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamAttributes"] = klass
	klass = &Class{
		"QXmlStreamEntityDeclaration",
		false,
		make([]*Class, 0),
		88,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamEntityDeclaration"] = klass
	klass = &Class{
		"QXmlStreamEntityResolver",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QXmlStreamEntityResolver"] = klass
	klass = &Class{
		"QXmlStreamNamespaceDeclaration",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamNamespaceDeclaration"] = klass
	klass = &Class{
		"QXmlStreamNotationDeclaration",
		false,
		make([]*Class, 0),
		56,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamNotationDeclaration"] = klass
	klass = &Class{
		"QXmlStreamReader",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QXmlStreamReader"] = klass
	klass = &Class{
		"QXmlStreamStringRef",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QXmlStreamStringRef"] = klass
	klass = &Class{
		"QXmlStreamWriter",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QXmlStreamWriter"] = klass
	classes["QAbstractAnimation"].parents = append(classes["QAbstractAnimation"].parents, classes["QObject"])
	classes["QAbstractEventDispatcher"].parents = append(classes["QAbstractEventDispatcher"].parents, classes["QObject"])
	classes["QAbstractFileEngine::MapExtensionOption"].parents = append(classes["QAbstractFileEngine::MapExtensionOption"].parents, classes["QAbstractFileEngine::ExtensionOption"])
	classes["QAbstractFileEngine::MapExtensionReturn"].parents = append(classes["QAbstractFileEngine::MapExtensionReturn"].parents, classes["QAbstractFileEngine::ExtensionReturn"])
	classes["QAbstractFileEngine::UnMapExtensionOption"].parents = append(classes["QAbstractFileEngine::UnMapExtensionOption"].parents, classes["QAbstractFileEngine::ExtensionOption"])
	classes["QAbstractItemModel"].parents = append(classes["QAbstractItemModel"].parents, classes["QObject"])
	classes["QAbstractListModel"].parents = append(classes["QAbstractListModel"].parents, classes["QAbstractItemModel"])
	classes["QAbstractState"].parents = append(classes["QAbstractState"].parents, classes["QObject"])
	classes["QAbstractTableModel"].parents = append(classes["QAbstractTableModel"].parents, classes["QAbstractItemModel"])
	classes["QAbstractTransition"].parents = append(classes["QAbstractTransition"].parents, classes["QObject"])
	classes["QAnimationGroup"].parents = append(classes["QAnimationGroup"].parents, classes["QAbstractAnimation"])
	classes["QAtomicInt"].parents = append(classes["QAtomicInt"].parents, classes["QBasicAtomicInt"])
	classes["QBuffer"].parents = append(classes["QBuffer"].parents, classes["QIODevice"])
	classes["QChildEvent"].parents = append(classes["QChildEvent"].parents, classes["QEvent"])
	classes["QCoreApplication"].parents = append(classes["QCoreApplication"].parents, classes["QObject"])
	classes["QDynamicPropertyChangeEvent"].parents = append(classes["QDynamicPropertyChangeEvent"].parents, classes["QEvent"])
	classes["QEventLoop"].parents = append(classes["QEventLoop"].parents, classes["QObject"])
	classes["QEventTransition"].parents = append(classes["QEventTransition"].parents, classes["QAbstractTransition"])
	classes["QFSFileEngine"].parents = append(classes["QFSFileEngine"].parents, classes["QAbstractFileEngine"])
	classes["QFile"].parents = append(classes["QFile"].parents, classes["QIODevice"])
	classes["QFileSystemWatcher"].parents = append(classes["QFileSystemWatcher"].parents, classes["QObject"])
	classes["QFinalState"].parents = append(classes["QFinalState"].parents, classes["QAbstractState"])
	classes["QFutureWatcherBase"].parents = append(classes["QFutureWatcherBase"].parents, classes["QObject"])
	classes["QGenericReturnArgument"].parents = append(classes["QGenericReturnArgument"].parents, classes["QGenericArgument"])
	classes["QHistoryState"].parents = append(classes["QHistoryState"].parents, classes["QAbstractState"])
	classes["QIODevice"].parents = append(classes["QIODevice"].parents, classes["QObject"])
	classes["QLibrary"].parents = append(classes["QLibrary"].parents, classes["QObject"])
	classes["QMimeData"].parents = append(classes["QMimeData"].parents, classes["QObject"])
	classes["QObjectCleanupHandler"].parents = append(classes["QObjectCleanupHandler"].parents, classes["QObject"])
	classes["QParallelAnimationGroup"].parents = append(classes["QParallelAnimationGroup"].parents, classes["QAnimationGroup"])
	classes["QPauseAnimation"].parents = append(classes["QPauseAnimation"].parents, classes["QAbstractAnimation"])
	classes["QPluginLoader"].parents = append(classes["QPluginLoader"].parents, classes["QObject"])
	classes["QProcess"].parents = append(classes["QProcess"].parents, classes["QIODevice"])
	classes["QPropertyAnimation"].parents = append(classes["QPropertyAnimation"].parents, classes["QVariantAnimation"])
	classes["QSequentialAnimationGroup"].parents = append(classes["QSequentialAnimationGroup"].parents, classes["QAnimationGroup"])
	classes["QSettings"].parents = append(classes["QSettings"].parents, classes["QObject"])
	classes["QSharedMemory"].parents = append(classes["QSharedMemory"].parents, classes["QObject"])
	classes["QSignalMapper"].parents = append(classes["QSignalMapper"].parents, classes["QObject"])
	classes["QSignalTransition"].parents = append(classes["QSignalTransition"].parents, classes["QAbstractTransition"])
	classes["QSocketNotifier"].parents = append(classes["QSocketNotifier"].parents, classes["QObject"])
	classes["QState"].parents = append(classes["QState"].parents, classes["QAbstractState"])
	classes["QStateMachine"].parents = append(classes["QStateMachine"].parents, classes["QState"])
	classes["QTemporaryFile"].parents = append(classes["QTemporaryFile"].parents, classes["QFile"])
	classes["QTextCodecFactoryInterface"].parents = append(classes["QTextCodecFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QTextCodecPlugin"].parents = append(classes["QTextCodecPlugin"].parents, classes["QObject"])
	classes["QTextCodecPlugin"].parents = append(classes["QTextCodecPlugin"].parents, classes["QTextCodecFactoryInterface"])
	classes["QTimeLine"].parents = append(classes["QTimeLine"].parents, classes["QObject"])
	classes["QTimer"].parents = append(classes["QTimer"].parents, classes["QObject"])
	classes["QTimerEvent"].parents = append(classes["QTimerEvent"].parents, classes["QEvent"])
	classes["QTranslator"].parents = append(classes["QTranslator"].parents, classes["QObject"])
	classes["QVariantAnimation"].parents = append(classes["QVariantAnimation"].parents, classes["QAbstractAnimation"])
	type_ = &Type{
		"FILE*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1] = type_
	type_ = &Type{
		"QAbstractAnimation*",
		classes["QAbstractAnimation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[2] = type_
	type_ = &Type{
		"QAbstractAnimation::DeletionPolicy",
		classes["QAbstractAnimation"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[3] = type_
	type_ = &Type{
		"QAbstractAnimation::Direction",
		classes["QAbstractAnimation"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[4] = type_
	type_ = &Type{
		"QAbstractAnimation::State",
		classes["QAbstractAnimation"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[5] = type_
	type_ = &Type{
		"QAbstractEventDispatcher*",
		classes["QAbstractEventDispatcher"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[6] = type_
	type_ = &Type{
		"QAbstractFileEngine*",
		classes["QAbstractFileEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[7] = type_
	type_ = &Type{
		"QAbstractFileEngine::Extension",
		classes["QAbstractFileEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[8] = type_
	type_ = &Type{
		"QAbstractFileEngine::ExtensionOption*",
		classes["QAbstractFileEngine::ExtensionOption"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[9] = type_
	type_ = &Type{
		"QAbstractFileEngine::ExtensionReturn*",
		classes["QAbstractFileEngine::ExtensionReturn"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[10] = type_
	type_ = &Type{
		"QAbstractFileEngine::FileFlag",
		classes["QAbstractFileEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[11] = type_
	type_ = &Type{
		"QAbstractFileEngine::FileName",
		classes["QAbstractFileEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[12] = type_
	type_ = &Type{
		"QAbstractFileEngine::FileOwner",
		classes["QAbstractFileEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[13] = type_
	type_ = &Type{
		"QAbstractFileEngine::FileTime",
		classes["QAbstractFileEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[14] = type_
	type_ = &Type{
		"QAbstractFileEngine::MapExtensionOption*",
		classes["QAbstractFileEngine::MapExtensionOption"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[15] = type_
	type_ = &Type{
		"QAbstractFileEngine::MapExtensionReturn*",
		classes["QAbstractFileEngine::MapExtensionReturn"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[16] = type_
	type_ = &Type{
		"QAbstractFileEngine::UnMapExtensionOption*",
		classes["QAbstractFileEngine::UnMapExtensionOption"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[17] = type_
	type_ = &Type{
		"QAbstractFileEngineHandler*",
		classes["QAbstractFileEngineHandler"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[18] = type_
	type_ = &Type{
		"QAbstractFileEngineIterator*",
		classes["QAbstractFileEngineIterator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[19] = type_
	type_ = &Type{
		"QAbstractFileEngineIterator::EntryInfoType",
		classes["QAbstractFileEngineIterator"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[20] = type_
	type_ = &Type{
		"QAbstractItemModel*",
		classes["QAbstractItemModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[21] = type_
	type_ = &Type{
		"QAbstractListModel*",
		classes["QAbstractListModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[22] = type_
	type_ = &Type{
		"QAbstractState*",
		classes["QAbstractState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[23] = type_
	type_ = &Type{
		"QAbstractTableModel*",
		classes["QAbstractTableModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[24] = type_
	type_ = &Type{
		"QAbstractTransition*",
		classes["QAbstractTransition"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[25] = type_
	type_ = &Type{
		"QAnimationGroup*",
		classes["QAnimationGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[26] = type_
	type_ = &Type{
		"QAtomicInt&",
		classes["QAtomicInt"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[27] = type_
	type_ = &Type{
		"QAtomicInt*",
		classes["QAtomicInt"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[28] = type_
	type_ = &Type{
		"QBasicAtomicInt&",
		classes["QBasicAtomicInt"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[29] = type_
	type_ = &Type{
		"QBasicAtomicInt*",
		classes["QBasicAtomicInt"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[30] = type_
	type_ = &Type{
		"QBasicTimer*",
		classes["QBasicTimer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[31] = type_
	type_ = &Type{
		"QBitArray",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[32] = type_
	type_ = &Type{
		"QBitArray&",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[33] = type_
	type_ = &Type{
		"QBitArray*",
		classes["QBitArray"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[34] = type_
	type_ = &Type{
		"QBitRef",
		classes["QBitRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[35] = type_
	type_ = &Type{
		"QBitRef&",
		classes["QBitRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[36] = type_
	type_ = &Type{
		"QBitRef*",
		classes["QBitRef"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[37] = type_
	type_ = &Type{
		"QBool",
		classes["QBool"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[38] = type_
	type_ = &Type{
		"QBool*",
		classes["QBool"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[39] = type_
	type_ = &Type{
		"QBuffer*",
		classes["QBuffer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[40] = type_
	type_ = &Type{
		"QByteArray",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[41] = type_
	type_ = &Type{
		"QByteArray&",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[42] = type_
	type_ = &Type{
		"QByteArray(*)(const QString&)",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[43] = type_
	type_ = &Type{
		"QByteArray*",
		classes["QByteArray"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[44] = type_
	type_ = &Type{
		"QByteArrayMatcher&",
		classes["QByteArrayMatcher"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[45] = type_
	type_ = &Type{
		"QByteArrayMatcher*",
		classes["QByteArrayMatcher"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[46] = type_
	type_ = &Type{
		"QByteRef",
		classes["QByteRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[47] = type_
	type_ = &Type{
		"QByteRef&",
		classes["QByteRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[48] = type_
	type_ = &Type{
		"QByteRef*",
		classes["QByteRef"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[49] = type_
	type_ = &Type{
		"QChar",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[50] = type_
	type_ = &Type{
		"QChar&",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[51] = type_
	type_ = &Type{
		"QChar*",
		classes["QChar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[52] = type_
	type_ = &Type{
		"QChar::Category",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[53] = type_
	type_ = &Type{
		"QChar::CombiningClass",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[54] = type_
	type_ = &Type{
		"QChar::Decomposition",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[55] = type_
	type_ = &Type{
		"QChar::Direction",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[56] = type_
	type_ = &Type{
		"QChar::Joining",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[57] = type_
	type_ = &Type{
		"QChar::SpecialCharacter",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[58] = type_
	type_ = &Type{
		"QChar::UnicodeVersion",
		classes["QChar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[59] = type_
	type_ = &Type{
		"QCharRef&",
		classes["QCharRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[60] = type_
	type_ = &Type{
		"QCharRef*",
		classes["QCharRef"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[61] = type_
	type_ = &Type{
		"QChildEvent*",
		classes["QChildEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[62] = type_
	type_ = &Type{
		"QCoreApplication*",
		classes["QCoreApplication"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[63] = type_
	type_ = &Type{
		"QCoreApplication::Encoding",
		classes["QCoreApplication"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[64] = type_
	type_ = &Type{
		"QCryptographicHash*",
		classes["QCryptographicHash"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[65] = type_
	type_ = &Type{
		"QCryptographicHash::Algorithm",
		classes["QCryptographicHash"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[66] = type_
	type_ = &Type{
		"QDataStream&",
		classes["QDataStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[67] = type_
	type_ = &Type{
		"QDataStream*",
		classes["QDataStream"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[68] = type_
	type_ = &Type{
		"QDataStream::ByteOrder",
		classes["QDataStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[69] = type_
	type_ = &Type{
		"QDataStream::FloatingPointPrecision",
		classes["QDataStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[70] = type_
	type_ = &Type{
		"QDataStream::Status",
		classes["QDataStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[71] = type_
	type_ = &Type{
		"QDataStream::Version",
		classes["QDataStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[72] = type_
	type_ = &Type{
		"QDate",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[73] = type_
	type_ = &Type{
		"QDate&",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[74] = type_
	type_ = &Type{
		"QDate*",
		classes["QDate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[75] = type_
	type_ = &Type{
		"QDate::MonthNameType",
		classes["QDate"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[76] = type_
	type_ = &Type{
		"QDateTime",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[77] = type_
	type_ = &Type{
		"QDateTime&",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[78] = type_
	type_ = &Type{
		"QDateTime*",
		classes["QDateTime"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[79] = type_
	type_ = &Type{
		"QDebug",
		classes["QDebug"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[80] = type_
	type_ = &Type{
		"QDebug&",
		classes["QDebug"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[81] = type_
	type_ = &Type{
		"QDebug*",
		classes["QDebug"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[82] = type_
	type_ = &Type{
		"QDir",
		classes["QDir"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[83] = type_
	type_ = &Type{
		"QDir&",
		classes["QDir"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[84] = type_
	type_ = &Type{
		"QDir*",
		classes["QDir"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[85] = type_
	type_ = &Type{
		"QDir::Filter",
		classes["QDir"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[86] = type_
	type_ = &Type{
		"QDir::SortFlag",
		classes["QDir"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[87] = type_
	type_ = &Type{
		"QDirIterator*",
		classes["QDirIterator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[88] = type_
	type_ = &Type{
		"QDirIterator::IteratorFlag",
		classes["QDirIterator"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[89] = type_
	type_ = &Type{
		"QDynamicPropertyChangeEvent*",
		classes["QDynamicPropertyChangeEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[90] = type_
	type_ = &Type{
		"QEasingCurve",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[91] = type_
	type_ = &Type{
		"QEasingCurve&",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[92] = type_
	type_ = &Type{
		"QEasingCurve*",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[93] = type_
	type_ = &Type{
		"QEasingCurve::Type",
		classes["QEasingCurve"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[94] = type_
	type_ = &Type{
		"QElapsedTimer*",
		classes["QElapsedTimer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[95] = type_
	type_ = &Type{
		"QElapsedTimer::ClockType",
		classes["QElapsedTimer"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[96] = type_
	type_ = &Type{
		"QEvent*",
		classes["QEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[97] = type_
	type_ = &Type{
		"QEvent::Type",
		classes["QEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[98] = type_
	type_ = &Type{
		"QEventLoop*",
		classes["QEventLoop"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[99] = type_
	type_ = &Type{
		"QEventLoop::ProcessEventsFlag",
		classes["QEventLoop"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[100] = type_
	type_ = &Type{
		"QEventPrivate*",
		classes["QEventPrivate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[101] = type_
	type_ = &Type{
		"QEventTransition*",
		classes["QEventTransition"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[102] = type_
	type_ = &Type{
		"QFSFileEngine*",
		classes["QFSFileEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[103] = type_
	type_ = &Type{
		"QFactoryInterface*",
		classes["QFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[104] = type_
	type_ = &Type{
		"QFile&",
		classes["QFile"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[105] = type_
	type_ = &Type{
		"QFile*",
		classes["QFile"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[106] = type_
	type_ = &Type{
		"QFile::FileError",
		classes["QFile"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[107] = type_
	type_ = &Type{
		"QFile::FileHandleFlag",
		classes["QFile"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[108] = type_
	type_ = &Type{
		"QFile::MemoryMapFlags",
		classes["QFile"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[109] = type_
	type_ = &Type{
		"QFile::Permission",
		classes["QFile"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[110] = type_
	type_ = &Type{
		"QFileInfo",
		classes["QFileInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[111] = type_
	type_ = &Type{
		"QFileInfo&",
		classes["QFileInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[112] = type_
	type_ = &Type{
		"QFileInfo*",
		classes["QFileInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[113] = type_
	type_ = &Type{
		"QFileSystemWatcher*",
		classes["QFileSystemWatcher"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[114] = type_
	type_ = &Type{
		"QFinalState*",
		classes["QFinalState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[115] = type_
	type_ = &Type{
		"QFlag*",
		classes["QFlag"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[116] = type_
	type_ = &Type{
		"QFlags<QAbstractFileEngine::FileFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[117] = type_
	type_ = &Type{
		"QFlags<QDir::Filter>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[118] = type_
	type_ = &Type{
		"QFlags<QDir::SortFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[119] = type_
	type_ = &Type{
		"QFlags<QDirIterator::IteratorFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[120] = type_
	type_ = &Type{
		"QFlags<QEventLoop::ProcessEventsFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[121] = type_
	type_ = &Type{
		"QFlags<QFile::FileHandleFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[122] = type_
	type_ = &Type{
		"QFlags<QFile::Permission>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[123] = type_
	type_ = &Type{
		"QFlags<QIODevice::OpenModeFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[124] = type_
	type_ = &Type{
		"QFlags<QLibrary::LoadHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[125] = type_
	type_ = &Type{
		"QFlags<QLocale::NumberOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[126] = type_
	type_ = &Type{
		"QFlags<QString::SectionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[127] = type_
	type_ = &Type{
		"QFlags<QTextBoundaryFinder::BoundaryReason>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[128] = type_
	type_ = &Type{
		"QFlags<QTextCodec::ConversionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[129] = type_
	type_ = &Type{
		"QFlags<QTextStream::NumberFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[130] = type_
	type_ = &Type{
		"QFlags<QUrl::FormattingOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[131] = type_
	type_ = &Type{
		"QFlags<Qt::AlignmentFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[132] = type_
	type_ = &Type{
		"QFlags<Qt::DockWidgetArea>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[133] = type_
	type_ = &Type{
		"QFlags<Qt::DropAction>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[134] = type_
	type_ = &Type{
		"QFlags<Qt::GestureFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[135] = type_
	type_ = &Type{
		"QFlags<Qt::ImageConversionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[136] = type_
	type_ = &Type{
		"QFlags<Qt::InputMethodHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[137] = type_
	type_ = &Type{
		"QFlags<Qt::ItemFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[138] = type_
	type_ = &Type{
		"QFlags<Qt::KeyboardModifier>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[139] = type_
	type_ = &Type{
		"QFlags<Qt::MatchFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[140] = type_
	type_ = &Type{
		"QFlags<Qt::MouseButton>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[141] = type_
	type_ = &Type{
		"QFlags<Qt::Orientation>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[142] = type_
	type_ = &Type{
		"QFlags<Qt::TextInteractionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[143] = type_
	type_ = &Type{
		"QFlags<Qt::ToolBarArea>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[144] = type_
	type_ = &Type{
		"QFlags<Qt::TouchPointState>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[145] = type_
	type_ = &Type{
		"QFlags<Qt::WindowState>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[146] = type_
	type_ = &Type{
		"QFlags<Qt::WindowType>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[147] = type_
	type_ = &Type{
		"QFutureInterfaceBase&",
		classes["QFutureInterfaceBase"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[148] = type_
	type_ = &Type{
		"QFutureInterfaceBase*",
		classes["QFutureInterfaceBase"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[149] = type_
	type_ = &Type{
		"QFutureInterfaceBase::State",
		classes["QFutureInterfaceBase"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[150] = type_
	type_ = &Type{
		"QGenericArgument",
		classes["QGenericArgument"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[151] = type_
	type_ = &Type{
		"QGenericArgument*",
		classes["QGenericArgument"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[152] = type_
	type_ = &Type{
		"QGenericReturnArgument",
		classes["QGenericReturnArgument"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[153] = type_
	type_ = &Type{
		"QGenericReturnArgument*",
		classes["QGenericReturnArgument"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[154] = type_
	type_ = &Type{
		"QHash<QString,QVariant>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[155] = type_
	type_ = &Type{
		"QHistoryState*",
		classes["QHistoryState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[156] = type_
	type_ = &Type{
		"QHistoryState::HistoryType",
		classes["QHistoryState"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[157] = type_
	type_ = &Type{
		"QIODevice*",
		classes["QIODevice"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[158] = type_
	type_ = &Type{
		"QIODevice::OpenModeFlag",
		classes["QIODevice"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[159] = type_
	type_ = &Type{
		"QIncompatibleFlag",
		classes["QIncompatibleFlag"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[160] = type_
	type_ = &Type{
		"QIncompatibleFlag*",
		classes["QIncompatibleFlag"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[161] = type_
	type_ = &Type{
		"QInternal*",
		classes["QInternal"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[162] = type_
	type_ = &Type{
		"QInternal::Callback",
		classes["QInternal"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[163] = type_
	type_ = &Type{
		"QInternal::DockPosition",
		classes["QInternal"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[164] = type_
	type_ = &Type{
		"QInternal::InternalFunction",
		classes["QInternal"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[165] = type_
	type_ = &Type{
		"QInternal::PaintDeviceFlags",
		classes["QInternal"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[166] = type_
	type_ = &Type{
		"QInternal::RelayoutType",
		classes["QInternal"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[167] = type_
	type_ = &Type{
		"QLatin1Char",
		classes["QLatin1Char"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[168] = type_
	type_ = &Type{
		"QLatin1Char*",
		classes["QLatin1Char"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[169] = type_
	type_ = &Type{
		"QLatin1String",
		classes["QLatin1String"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[170] = type_
	type_ = &Type{
		"QLatin1String&",
		classes["QLatin1String"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[171] = type_
	type_ = &Type{
		"QLatin1String*",
		classes["QLatin1String"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[172] = type_
	type_ = &Type{
		"QLibrary*",
		classes["QLibrary"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[173] = type_
	type_ = &Type{
		"QLibrary::LoadHint",
		classes["QLibrary"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[174] = type_
	type_ = &Type{
		"QLibraryInfo*",
		classes["QLibraryInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[175] = type_
	type_ = &Type{
		"QLibraryInfo::LibraryLocation",
		classes["QLibraryInfo"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[176] = type_
	type_ = &Type{
		"QLine",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[177] = type_
	type_ = &Type{
		"QLine&",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[178] = type_
	type_ = &Type{
		"QLine*",
		classes["QLine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[179] = type_
	type_ = &Type{
		"QLineF",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[180] = type_
	type_ = &Type{
		"QLineF&",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[181] = type_
	type_ = &Type{
		"QLineF*",
		classes["QLineF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[182] = type_
	type_ = &Type{
		"QLineF::IntersectType",
		classes["QLineF"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[183] = type_
	type_ = &Type{
		"QList<QAbstractAnimation*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[184] = type_
	type_ = &Type{
		"QList<QAbstractState*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[185] = type_
	type_ = &Type{
		"QList<QAbstractTransition*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[186] = type_
	type_ = &Type{
		"QList<QByteArray>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[187] = type_
	type_ = &Type{
		"QList<QFileInfo>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[188] = type_
	type_ = &Type{
		"QList<QLocale::Country>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[189] = type_
	type_ = &Type{
		"QList<QLocale>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[190] = type_
	type_ = &Type{
		"QList<QModelIndex>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[191] = type_
	type_ = &Type{
		"QList<QObject*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[192] = type_
	type_ = &Type{
		"QList<QPair<QByteArray,QByteArray> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[193] = type_
	type_ = &Type{
		"QList<QPair<QString,QString> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[194] = type_
	type_ = &Type{
		"QList<QPair<int,int> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[195] = type_
	type_ = &Type{
		"QList<QUrl>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[196] = type_
	type_ = &Type{
		"QList<QVariant>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[197] = type_
	type_ = &Type{
		"QList<Qt::DayOfWeek>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[198] = type_
	type_ = &Type{
		"QList<int>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[199] = type_
	type_ = &Type{
		"QList<void*>*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[200] = type_
	type_ = &Type{
		"QLocale",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[201] = type_
	type_ = &Type{
		"QLocale&",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[202] = type_
	type_ = &Type{
		"QLocale*",
		classes["QLocale"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[203] = type_
	type_ = &Type{
		"QLocale::Country",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[204] = type_
	type_ = &Type{
		"QLocale::CurrencySymbolFormat",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[205] = type_
	type_ = &Type{
		"QLocale::FormatType",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[206] = type_
	type_ = &Type{
		"QLocale::Language",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[207] = type_
	type_ = &Type{
		"QLocale::MeasurementSystem",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[208] = type_
	type_ = &Type{
		"QLocale::NumberOption",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[209] = type_
	type_ = &Type{
		"QLocale::QuotationStyle",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[210] = type_
	type_ = &Type{
		"QLocale::Script",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[211] = type_
	type_ = &Type{
		"QMap<QString,QVariant>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[212] = type_
	type_ = &Type{
		"QMap<int,QVariant>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[213] = type_
	type_ = &Type{
		"QMargins*",
		classes["QMargins"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[214] = type_
	type_ = &Type{
		"QMetaClassInfo",
		classes["QMetaClassInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[215] = type_
	type_ = &Type{
		"QMetaClassInfo*",
		classes["QMetaClassInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[216] = type_
	type_ = &Type{
		"QMetaEnum",
		classes["QMetaEnum"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[217] = type_
	type_ = &Type{
		"QMetaEnum*",
		classes["QMetaEnum"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[218] = type_
	type_ = &Type{
		"QMetaMethod",
		classes["QMetaMethod"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[219] = type_
	type_ = &Type{
		"QMetaMethod*",
		classes["QMetaMethod"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[220] = type_
	type_ = &Type{
		"QMetaMethod::Access",
		classes["QMetaMethod"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[221] = type_
	type_ = &Type{
		"QMetaMethod::Attributes",
		classes["QMetaMethod"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[222] = type_
	type_ = &Type{
		"QMetaMethod::MethodType",
		classes["QMetaMethod"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[223] = type_
	type_ = &Type{
		"QMetaObject*",
		classes["QMetaObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[224] = type_
	type_ = &Type{
		"QMetaObject::Call",
		classes["QMetaObject"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[225] = type_
	type_ = &Type{
		"QMetaProperty",
		classes["QMetaProperty"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[226] = type_
	type_ = &Type{
		"QMetaProperty*",
		classes["QMetaProperty"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[227] = type_
	type_ = &Type{
		"QMetaType*",
		classes["QMetaType"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[228] = type_
	type_ = &Type{
		"QMetaType::Type",
		classes["QMetaType"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[229] = type_
	type_ = &Type{
		"QMimeData*",
		classes["QMimeData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[230] = type_
	type_ = &Type{
		"QModelIndex",
		classes["QModelIndex"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[231] = type_
	type_ = &Type{
		"QModelIndex*",
		classes["QModelIndex"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[232] = type_
	type_ = &Type{
		"QMutex*",
		classes["QMutex"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[233] = type_
	type_ = &Type{
		"QNoDebug&",
		classes["QNoDebug"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[234] = type_
	type_ = &Type{
		"QNoDebug*",
		classes["QNoDebug"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[235] = type_
	type_ = &Type{
		"QObject*",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[236] = type_
	type_ = &Type{
		"QObject*(*)()",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[237] = type_
	type_ = &Type{
		"QObject**",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[238] = type_
	type_ = &Type{
		"QObjectCleanupHandler*",
		classes["QObjectCleanupHandler"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[239] = type_
	type_ = &Type{
		"QObjectUserData*",
		classes["QObjectUserData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[240] = type_
	type_ = &Type{
		"QParallelAnimationGroup*",
		classes["QParallelAnimationGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[241] = type_
	type_ = &Type{
		"QPauseAnimation*",
		classes["QPauseAnimation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[242] = type_
	type_ = &Type{
		"QPersistentModelIndex&",
		classes["QPersistentModelIndex"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[243] = type_
	type_ = &Type{
		"QPersistentModelIndex*",
		classes["QPersistentModelIndex"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[244] = type_
	type_ = &Type{
		"QPluginLoader*",
		classes["QPluginLoader"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[245] = type_
	type_ = &Type{
		"QPoint",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[246] = type_
	type_ = &Type{
		"QPoint&",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[247] = type_
	type_ = &Type{
		"QPoint*",
		classes["QPoint"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[248] = type_
	type_ = &Type{
		"QPointF",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[249] = type_
	type_ = &Type{
		"QPointF&",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[250] = type_
	type_ = &Type{
		"QPointF*",
		classes["QPointF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[251] = type_
	type_ = &Type{
		"QPostEventList*",
		classes["QPostEventList"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[252] = type_
	type_ = &Type{
		"QProcess*",
		classes["QProcess"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[253] = type_
	type_ = &Type{
		"QProcess::ExitStatus",
		classes["QProcess"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[254] = type_
	type_ = &Type{
		"QProcess::ProcessChannel",
		classes["QProcess"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[255] = type_
	type_ = &Type{
		"QProcess::ProcessChannelMode",
		classes["QProcess"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[256] = type_
	type_ = &Type{
		"QProcess::ProcessError",
		classes["QProcess"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[257] = type_
	type_ = &Type{
		"QProcess::ProcessState",
		classes["QProcess"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[258] = type_
	type_ = &Type{
		"QProcessEnvironment",
		classes["QProcessEnvironment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[259] = type_
	type_ = &Type{
		"QProcessEnvironment&",
		classes["QProcessEnvironment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[260] = type_
	type_ = &Type{
		"QProcessEnvironment*",
		classes["QProcessEnvironment"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[261] = type_
	type_ = &Type{
		"QPropertyAnimation*",
		classes["QPropertyAnimation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[262] = type_
	type_ = &Type{
		"QReadLocker*",
		classes["QReadLocker"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[263] = type_
	type_ = &Type{
		"QReadWriteLock*",
		classes["QReadWriteLock"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[264] = type_
	type_ = &Type{
		"QReadWriteLock::RecursionMode",
		classes["QReadWriteLock"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[265] = type_
	type_ = &Type{
		"QRect",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[266] = type_
	type_ = &Type{
		"QRect&",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[267] = type_
	type_ = &Type{
		"QRect*",
		classes["QRect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[268] = type_
	type_ = &Type{
		"QRectF",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[269] = type_
	type_ = &Type{
		"QRectF&",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[270] = type_
	type_ = &Type{
		"QRectF*",
		classes["QRectF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[271] = type_
	type_ = &Type{
		"QRegExp",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[272] = type_
	type_ = &Type{
		"QRegExp&",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[273] = type_
	type_ = &Type{
		"QRegExp*",
		classes["QRegExp"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[274] = type_
	type_ = &Type{
		"QRegExp::CaretMode",
		classes["QRegExp"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[275] = type_
	type_ = &Type{
		"QRegExp::PatternSyntax",
		classes["QRegExp"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[276] = type_
	type_ = &Type{
		"QResource*",
		classes["QResource"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[277] = type_
	type_ = &Type{
		"QRunnable*",
		classes["QRunnable"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[278] = type_
	type_ = &Type{
		"QSemaphore*",
		classes["QSemaphore"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[279] = type_
	type_ = &Type{
		"QSequentialAnimationGroup*",
		classes["QSequentialAnimationGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[280] = type_
	type_ = &Type{
		"QSet<QAbstractState*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[281] = type_
	type_ = &Type{
		"QSettings*",
		classes["QSettings"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[282] = type_
	type_ = &Type{
		"QSettings::Format",
		classes["QSettings"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[283] = type_
	type_ = &Type{
		"QSettings::Scope",
		classes["QSettings"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[284] = type_
	type_ = &Type{
		"QSettings::Status",
		classes["QSettings"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[285] = type_
	type_ = &Type{
		"QSharedData*",
		classes["QSharedData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[286] = type_
	type_ = &Type{
		"QSharedMemory*",
		classes["QSharedMemory"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[287] = type_
	type_ = &Type{
		"QSharedMemory::AccessMode",
		classes["QSharedMemory"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[288] = type_
	type_ = &Type{
		"QSharedMemory::SharedMemoryError",
		classes["QSharedMemory"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[289] = type_
	type_ = &Type{
		"QSignalMapper*",
		classes["QSignalMapper"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[290] = type_
	type_ = &Type{
		"QSignalTransition*",
		classes["QSignalTransition"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[291] = type_
	type_ = &Type{
		"QSize",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[292] = type_
	type_ = &Type{
		"QSize&",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[293] = type_
	type_ = &Type{
		"QSize*",
		classes["QSize"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[294] = type_
	type_ = &Type{
		"QSizeF",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[295] = type_
	type_ = &Type{
		"QSizeF&",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[296] = type_
	type_ = &Type{
		"QSizeF*",
		classes["QSizeF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[297] = type_
	type_ = &Type{
		"QSocketNotifier*",
		classes["QSocketNotifier"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[298] = type_
	type_ = &Type{
		"QSocketNotifier::Type",
		classes["QSocketNotifier"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[299] = type_
	type_ = &Type{
		"QState*",
		classes["QState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[300] = type_
	type_ = &Type{
		"QState::ChildMode",
		classes["QState"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[301] = type_
	type_ = &Type{
		"QStateMachine*",
		classes["QStateMachine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[302] = type_
	type_ = &Type{
		"QStateMachine::Error",
		classes["QStateMachine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[303] = type_
	type_ = &Type{
		"QStateMachine::EventPriority",
		classes["QStateMachine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[304] = type_
	type_ = &Type{
		"QStateMachine::RestorePolicy",
		classes["QStateMachine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[305] = type_
	type_ = &Type{
		"QString",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[306] = type_
	type_ = &Type{
		"QString&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[307] = type_
	type_ = &Type{
		"QString(*)(const QByteArray&)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[308] = type_
	type_ = &Type{
		"QString*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[309] = type_
	type_ = &Type{
		"QString::Null",
		classes["QString::Null"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[310] = type_
	type_ = &Type{
		"QString::SectionFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[311] = type_
	type_ = &Type{
		"QStringList",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[312] = type_
	type_ = &Type{
		"QStringList&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[313] = type_
	type_ = &Type{
		"QStringList*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[314] = type_
	type_ = &Type{
		"QStringMatcher&",
		classes["QStringMatcher"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[315] = type_
	type_ = &Type{
		"QStringMatcher*",
		classes["QStringMatcher"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[316] = type_
	type_ = &Type{
		"QStringRef",
		classes["QStringRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[317] = type_
	type_ = &Type{
		"QStringRef&",
		classes["QStringRef"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[318] = type_
	type_ = &Type{
		"QStringRef*",
		classes["QStringRef"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[319] = type_
	type_ = &Type{
		"QSysInfo*",
		classes["QSysInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[320] = type_
	type_ = &Type{
		"QSysInfo::Endian",
		classes["QSysInfo"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[321] = type_
	type_ = &Type{
		"QSysInfo::Sizes",
		classes["QSysInfo"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[322] = type_
	type_ = &Type{
		"QSystemLocale*",
		classes["QSystemLocale"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[323] = type_
	type_ = &Type{
		"QSystemLocale::QueryType",
		classes["QSystemLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[324] = type_
	type_ = &Type{
		"QSystemSemaphore*",
		classes["QSystemSemaphore"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[325] = type_
	type_ = &Type{
		"QSystemSemaphore::AccessMode",
		classes["QSystemSemaphore"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[326] = type_
	type_ = &Type{
		"QSystemSemaphore::SystemSemaphoreError",
		classes["QSystemSemaphore"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[327] = type_
	type_ = &Type{
		"QTemporaryFile*",
		classes["QTemporaryFile"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[328] = type_
	type_ = &Type{
		"QTextBoundaryFinder&",
		classes["QTextBoundaryFinder"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[329] = type_
	type_ = &Type{
		"QTextBoundaryFinder*",
		classes["QTextBoundaryFinder"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[330] = type_
	type_ = &Type{
		"QTextBoundaryFinder::BoundaryReason",
		classes["QTextBoundaryFinder"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[331] = type_
	type_ = &Type{
		"QTextBoundaryFinder::BoundaryType",
		classes["QTextBoundaryFinder"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[332] = type_
	type_ = &Type{
		"QTextCodec*",
		classes["QTextCodec"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[333] = type_
	type_ = &Type{
		"QTextCodec::ConversionFlag",
		classes["QTextCodec"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[334] = type_
	type_ = &Type{
		"QTextCodec::ConverterState*",
		classes["QTextCodec::ConverterState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[335] = type_
	type_ = &Type{
		"QTextCodecFactoryInterface*",
		classes["QTextCodecFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[336] = type_
	type_ = &Type{
		"QTextCodecPlugin*",
		classes["QTextCodecPlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[337] = type_
	type_ = &Type{
		"QTextDecoder*",
		classes["QTextDecoder"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[338] = type_
	type_ = &Type{
		"QTextEncoder*",
		classes["QTextEncoder"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[339] = type_
	type_ = &Type{
		"QTextStream&",
		classes["QTextStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[340] = type_
	type_ = &Type{
		"QTextStream&(*)(QTextStream&)",
		classes["QTextStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[341] = type_
	type_ = &Type{
		"QTextStream*",
		classes["QTextStream"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[342] = type_
	type_ = &Type{
		"QTextStream::FieldAlignment",
		classes["QTextStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[343] = type_
	type_ = &Type{
		"QTextStream::NumberFlag",
		classes["QTextStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[344] = type_
	type_ = &Type{
		"QTextStream::RealNumberNotation",
		classes["QTextStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[345] = type_
	type_ = &Type{
		"QTextStream::Status",
		classes["QTextStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[346] = type_
	type_ = &Type{
		"QTextStreamManipulator",
		classes["QTextStreamManipulator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[347] = type_
	type_ = &Type{
		"QThread*",
		classes["QThread"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[348] = type_
	type_ = &Type{
		"QTime",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[349] = type_
	type_ = &Type{
		"QTime&",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[350] = type_
	type_ = &Type{
		"QTime*",
		classes["QTime"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[351] = type_
	type_ = &Type{
		"QTime::TimeFlag",
		classes["QTime"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[352] = type_
	type_ = &Type{
		"QTimeLine*",
		classes["QTimeLine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[353] = type_
	type_ = &Type{
		"QTimeLine::CurveShape",
		classes["QTimeLine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[354] = type_
	type_ = &Type{
		"QTimeLine::Direction",
		classes["QTimeLine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[355] = type_
	type_ = &Type{
		"QTimeLine::State",
		classes["QTimeLine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[356] = type_
	type_ = &Type{
		"QTimer*",
		classes["QTimer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[357] = type_
	type_ = &Type{
		"QTimerEvent*",
		classes["QTimerEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[358] = type_
	type_ = &Type{
		"QTranslator*",
		classes["QTranslator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[359] = type_
	type_ = &Type{
		"QUrl",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[360] = type_
	type_ = &Type{
		"QUrl&",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[361] = type_
	type_ = &Type{
		"QUrl*",
		classes["QUrl"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[362] = type_
	type_ = &Type{
		"QUrl::FormattingOption",
		classes["QUrl"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[363] = type_
	type_ = &Type{
		"QUrl::ParsingMode",
		classes["QUrl"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[364] = type_
	type_ = &Type{
		"QUrlPrivate*&",
		classes["QUrlPrivate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[365] = type_
	type_ = &Type{
		"QUuid",
		classes["QUuid"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[366] = type_
	type_ = &Type{
		"QUuid&",
		classes["QUuid"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[367] = type_
	type_ = &Type{
		"QUuid*",
		classes["QUuid"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[368] = type_
	type_ = &Type{
		"QUuid::Variant",
		classes["QUuid"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[369] = type_
	type_ = &Type{
		"QUuid::Version",
		classes["QUuid"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[370] = type_
	type_ = &Type{
		"QVariant",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[371] = type_
	type_ = &Type{
		"QVariant&",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[372] = type_
	type_ = &Type{
		"QVariant*",
		classes["QVariant"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[373] = type_
	type_ = &Type{
		"QVariant::Private&",
		classes["QVariant::Private"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[374] = type_
	type_ = &Type{
		"QVariant::Type",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[375] = type_
	type_ = &Type{
		"QVariant::Type&",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[376] = type_
	type_ = &Type{
		"QVariantAnimation*",
		classes["QVariantAnimation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[377] = type_
	type_ = &Type{
		"QVariantComparisonHelper*",
		classes["QVariantComparisonHelper"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[378] = type_
	type_ = &Type{
		"QVector<QPair<double,QVariant> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[379] = type_
	type_ = &Type{
		"QVector<QXmlStreamEntityDeclaration>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[380] = type_
	type_ = &Type{
		"QVector<QXmlStreamNamespaceDeclaration>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[381] = type_
	type_ = &Type{
		"QVector<QXmlStreamNotationDeclaration>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[382] = type_
	type_ = &Type{
		"QVector<unsigned int>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[383] = type_
	type_ = &Type{
		"QWidget*",
		classes["QWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[384] = type_
	type_ = &Type{
		"QWriteLocker*",
		classes["QWriteLocker"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[385] = type_
	type_ = &Type{
		"QXmlStreamAttribute&",
		classes["QXmlStreamAttribute"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[386] = type_
	type_ = &Type{
		"QXmlStreamAttribute*",
		classes["QXmlStreamAttribute"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[387] = type_
	type_ = &Type{
		"QXmlStreamAttributes",
		classes["QXmlStreamAttributes"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[388] = type_
	type_ = &Type{
		"QXmlStreamAttributes*",
		classes["QXmlStreamAttributes"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[389] = type_
	type_ = &Type{
		"QXmlStreamEntityDeclaration&",
		classes["QXmlStreamEntityDeclaration"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[390] = type_
	type_ = &Type{
		"QXmlStreamEntityDeclaration*",
		classes["QXmlStreamEntityDeclaration"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[391] = type_
	type_ = &Type{
		"QXmlStreamEntityResolver*",
		classes["QXmlStreamEntityResolver"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[392] = type_
	type_ = &Type{
		"QXmlStreamNamespaceDeclaration&",
		classes["QXmlStreamNamespaceDeclaration"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[393] = type_
	type_ = &Type{
		"QXmlStreamNamespaceDeclaration*",
		classes["QXmlStreamNamespaceDeclaration"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[394] = type_
	type_ = &Type{
		"QXmlStreamNotationDeclaration&",
		classes["QXmlStreamNotationDeclaration"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[395] = type_
	type_ = &Type{
		"QXmlStreamNotationDeclaration*",
		classes["QXmlStreamNotationDeclaration"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[396] = type_
	type_ = &Type{
		"QXmlStreamReader*",
		classes["QXmlStreamReader"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[397] = type_
	type_ = &Type{
		"QXmlStreamReader::Error",
		classes["QXmlStreamReader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[398] = type_
	type_ = &Type{
		"QXmlStreamReader::ReadElementTextBehaviour",
		classes["QXmlStreamReader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[399] = type_
	type_ = &Type{
		"QXmlStreamReader::TokenType",
		classes["QXmlStreamReader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[400] = type_
	type_ = &Type{
		"QXmlStreamStringRef*",
		classes["QXmlStreamStringRef"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[401] = type_
	type_ = &Type{
		"QXmlStreamWriter*",
		classes["QXmlStreamWriter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[402] = type_
	type_ = &Type{
		"Qt::AlignmentFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[403] = type_
	type_ = &Type{
		"Qt::AnchorAttribute",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[404] = type_
	type_ = &Type{
		"Qt::AnchorPoint",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[405] = type_
	type_ = &Type{
		"Qt::ApplicationAttribute",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[406] = type_
	type_ = &Type{
		"Qt::ArrowType",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[407] = type_
	type_ = &Type{
		"Qt::AspectRatioMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[408] = type_
	type_ = &Type{
		"Qt::Axis",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[409] = type_
	type_ = &Type{
		"Qt::BGMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[410] = type_
	type_ = &Type{
		"Qt::BrushStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[411] = type_
	type_ = &Type{
		"Qt::CaseSensitivity",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[412] = type_
	type_ = &Type{
		"Qt::CheckState",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[413] = type_
	type_ = &Type{
		"Qt::ClipOperation",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[414] = type_
	type_ = &Type{
		"Qt::ConnectionType",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[415] = type_
	type_ = &Type{
		"Qt::ContextMenuPolicy",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[416] = type_
	type_ = &Type{
		"Qt::CoordinateSystem",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[417] = type_
	type_ = &Type{
		"Qt::Corner",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[418] = type_
	type_ = &Type{
		"Qt::CursorMoveStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[419] = type_
	type_ = &Type{
		"Qt::CursorShape",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[420] = type_
	type_ = &Type{
		"Qt::DateFormat",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[421] = type_
	type_ = &Type{
		"Qt::DayOfWeek",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[422] = type_
	type_ = &Type{
		"Qt::DockWidgetArea",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[423] = type_
	type_ = &Type{
		"Qt::DockWidgetAreaSizes",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[424] = type_
	type_ = &Type{
		"Qt::DropAction",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[425] = type_
	type_ = &Type{
		"Qt::EventPriority",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[426] = type_
	type_ = &Type{
		"Qt::FillRule",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[427] = type_
	type_ = &Type{
		"Qt::FocusPolicy",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[428] = type_
	type_ = &Type{
		"Qt::FocusReason",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[429] = type_
	type_ = &Type{
		"Qt::GestureFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[430] = type_
	type_ = &Type{
		"Qt::GestureState",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[431] = type_
	type_ = &Type{
		"Qt::GestureType",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[432] = type_
	type_ = &Type{
		"Qt::GlobalColor",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[433] = type_
	type_ = &Type{
		"Qt::ImageConversionFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[434] = type_
	type_ = &Type{
		"Qt::Initialization",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[435] = type_
	type_ = &Type{
		"Qt::InputMethodHint",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[436] = type_
	type_ = &Type{
		"Qt::InputMethodQuery",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[437] = type_
	type_ = &Type{
		"Qt::ItemDataRole",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[438] = type_
	type_ = &Type{
		"Qt::ItemFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[439] = type_
	type_ = &Type{
		"Qt::ItemSelectionMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[440] = type_
	type_ = &Type{
		"Qt::Key",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[441] = type_
	type_ = &Type{
		"Qt::KeyboardModifier",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[442] = type_
	type_ = &Type{
		"Qt::LayoutDirection",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[443] = type_
	type_ = &Type{
		"Qt::MaskMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[444] = type_
	type_ = &Type{
		"Qt::MatchFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[445] = type_
	type_ = &Type{
		"Qt::Modifier",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[446] = type_
	type_ = &Type{
		"Qt::MouseButton",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[447] = type_
	type_ = &Type{
		"Qt::NavigationMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[448] = type_
	type_ = &Type{
		"Qt::Orientation",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[449] = type_
	type_ = &Type{
		"Qt::PenCapStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[450] = type_
	type_ = &Type{
		"Qt::PenJoinStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[451] = type_
	type_ = &Type{
		"Qt::PenStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[452] = type_
	type_ = &Type{
		"Qt::ScrollBarPolicy",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[453] = type_
	type_ = &Type{
		"Qt::ShortcutContext",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[454] = type_
	type_ = &Type{
		"Qt::SizeHint",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[455] = type_
	type_ = &Type{
		"Qt::SizeMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[456] = type_
	type_ = &Type{
		"Qt::SortOrder",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[457] = type_
	type_ = &Type{
		"Qt::TextElideMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[458] = type_
	type_ = &Type{
		"Qt::TextFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[459] = type_
	type_ = &Type{
		"Qt::TextFormat",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[460] = type_
	type_ = &Type{
		"Qt::TextInteractionFlag",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[461] = type_
	type_ = &Type{
		"Qt::TileRule",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[462] = type_
	type_ = &Type{
		"Qt::TimeSpec",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[463] = type_
	type_ = &Type{
		"Qt::ToolBarArea",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[464] = type_
	type_ = &Type{
		"Qt::ToolBarAreaSizes",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[465] = type_
	type_ = &Type{
		"Qt::ToolButtonStyle",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[466] = type_
	type_ = &Type{
		"Qt::TouchPointState",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[467] = type_
	type_ = &Type{
		"Qt::TransformationMode",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[468] = type_
	type_ = &Type{
		"Qt::UIEffect",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[469] = type_
	type_ = &Type{
		"Qt::WidgetAttribute",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[470] = type_
	type_ = &Type{
		"Qt::WindowFrameSection",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[471] = type_
	type_ = &Type{
		"Qt::WindowModality",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[472] = type_
	type_ = &Type{
		"Qt::WindowState",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[473] = type_
	type_ = &Type{
		"Qt::WindowType",
		classes["Qt"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[474] = type_
	type_ = &Type{
		"QtConcurrent::ReduceOption",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[475] = type_
	type_ = &Type{
		"QtConcurrent::ThreadFunctionResult",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[476] = type_
	type_ = &Type{
		"QtMsgType",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[477] = type_
	type_ = &Type{
		"QtValidLicenseForActiveQtModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[478] = type_
	type_ = &Type{
		"QtValidLicenseForCoreModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[479] = type_
	type_ = &Type{
		"QtValidLicenseForDBusModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[480] = type_
	type_ = &Type{
		"QtValidLicenseForDeclarativeModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[481] = type_
	type_ = &Type{
		"QtValidLicenseForGuiModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[482] = type_
	type_ = &Type{
		"QtValidLicenseForHelpModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[483] = type_
	type_ = &Type{
		"QtValidLicenseForMultimediaModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[484] = type_
	type_ = &Type{
		"QtValidLicenseForNetworkModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[485] = type_
	type_ = &Type{
		"QtValidLicenseForOpenGLModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[486] = type_
	type_ = &Type{
		"QtValidLicenseForOpenVGModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[487] = type_
	type_ = &Type{
		"QtValidLicenseForQt3SupportLightModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[488] = type_
	type_ = &Type{
		"QtValidLicenseForQt3SupportModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[489] = type_
	type_ = &Type{
		"QtValidLicenseForScriptModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[490] = type_
	type_ = &Type{
		"QtValidLicenseForScriptToolsModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[491] = type_
	type_ = &Type{
		"QtValidLicenseForSqlModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[492] = type_
	type_ = &Type{
		"QtValidLicenseForSvgModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[493] = type_
	type_ = &Type{
		"QtValidLicenseForTestModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[494] = type_
	type_ = &Type{
		"QtValidLicenseForXmlModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[495] = type_
	type_ = &Type{
		"QtValidLicenseForXmlPatternsModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[496] = type_
	type_ = &Type{
		"__codecvt_result",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[497] = type_
	type_ = &Type{
		"bool",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[498] = type_
	type_ = &Type{
		"bool&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[499] = type_
	type_ = &Type{
		"bool(*)(QIODevice&,QMap<QString,QVariant>&)",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[500] = type_
	type_ = &Type{
		"bool(*)(QIODevice&,const QMap<QString,QVariant>&)",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[501] = type_
	type_ = &Type{
		"bool(*)(void*)",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[502] = type_
	type_ = &Type{
		"bool(*)(void**)",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[503] = type_
	type_ = &Type{
		"bool(*)(void*,long*)",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[504] = type_
	type_ = &Type{
		"bool*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[505] = type_
	type_ = &Type{
		"char",
		nil,
		T_CHAR,
		KIND_STACK,
		false,
	}
	types[506] = type_
	type_ = &Type{
		"char&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[507] = type_
	type_ = &Type{
		"char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[508] = type_
	type_ = &Type{
		"char*&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[509] = type_
	type_ = &Type{
		"char**",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[510] = type_
	type_ = &Type{
		"const QAbstractFileEngine::ExtensionOption&",
		classes["QAbstractFileEngine::ExtensionOption"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[511] = type_
	type_ = &Type{
		"const QAbstractFileEngine::ExtensionOption*",
		classes["QAbstractFileEngine::ExtensionOption"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[512] = type_
	type_ = &Type{
		"const QAbstractFileEngine::ExtensionReturn&",
		classes["QAbstractFileEngine::ExtensionReturn"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[513] = type_
	type_ = &Type{
		"const QAbstractFileEngine::MapExtensionOption&",
		classes["QAbstractFileEngine::MapExtensionOption"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[514] = type_
	type_ = &Type{
		"const QAbstractFileEngine::MapExtensionReturn&",
		classes["QAbstractFileEngine::MapExtensionReturn"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[515] = type_
	type_ = &Type{
		"const QAbstractFileEngine::UnMapExtensionOption&",
		classes["QAbstractFileEngine::UnMapExtensionOption"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[516] = type_
	type_ = &Type{
		"const QAbstractFileEngineHandler&",
		classes["QAbstractFileEngineHandler"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[517] = type_
	type_ = &Type{
		"const QAbstractItemModel*",
		classes["QAbstractItemModel"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[518] = type_
	type_ = &Type{
		"const QAtomicInt&",
		classes["QAtomicInt"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[519] = type_
	type_ = &Type{
		"const QBasicAtomicInt&",
		classes["QBasicAtomicInt"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[520] = type_
	type_ = &Type{
		"const QBasicTimer&",
		classes["QBasicTimer"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[521] = type_
	type_ = &Type{
		"const QBitArray&",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[522] = type_
	type_ = &Type{
		"const QBitRef&",
		classes["QBitRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[523] = type_
	type_ = &Type{
		"const QBool&",
		classes["QBool"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[524] = type_
	type_ = &Type{
		"const QByteArray",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[525] = type_
	type_ = &Type{
		"const QByteArray&",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[526] = type_
	type_ = &Type{
		"const QByteArrayMatcher&",
		classes["QByteArrayMatcher"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[527] = type_
	type_ = &Type{
		"const QByteRef&",
		classes["QByteRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[528] = type_
	type_ = &Type{
		"const QChar",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[529] = type_
	type_ = &Type{
		"const QChar&",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[530] = type_
	type_ = &Type{
		"const QChar*",
		classes["QChar"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[531] = type_
	type_ = &Type{
		"const QCharRef&",
		classes["QCharRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[532] = type_
	type_ = &Type{
		"const QChildEvent&",
		classes["QChildEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[533] = type_
	type_ = &Type{
		"const QDate&",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[534] = type_
	type_ = &Type{
		"const QDateTime&",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[535] = type_
	type_ = &Type{
		"const QDebug&",
		classes["QDebug"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[536] = type_
	type_ = &Type{
		"const QDir&",
		classes["QDir"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[537] = type_
	type_ = &Type{
		"const QDynamicPropertyChangeEvent&",
		classes["QDynamicPropertyChangeEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[538] = type_
	type_ = &Type{
		"const QEasingCurve&",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[539] = type_
	type_ = &Type{
		"const QElapsedTimer&",
		classes["QElapsedTimer"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[540] = type_
	type_ = &Type{
		"const QEvent&",
		classes["QEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[541] = type_
	type_ = &Type{
		"const QFactoryInterface&",
		classes["QFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[542] = type_
	type_ = &Type{
		"const QFile&",
		classes["QFile"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[543] = type_
	type_ = &Type{
		"const QFileInfo&",
		classes["QFileInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[544] = type_
	type_ = &Type{
		"const QFlag&",
		classes["QFlag"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[545] = type_
	type_ = &Type{
		"const QFutureInterfaceBase&",
		classes["QFutureInterfaceBase"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[546] = type_
	type_ = &Type{
		"const QGenericArgument&",
		classes["QGenericArgument"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[547] = type_
	type_ = &Type{
		"const QGenericReturnArgument&",
		classes["QGenericReturnArgument"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[548] = type_
	type_ = &Type{
		"const QHash<QString,QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[549] = type_
	type_ = &Type{
		"const QHash<int,QByteArray>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[550] = type_
	type_ = &Type{
		"const QHashDummyValue&",
		classes["QHashDummyValue"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[551] = type_
	type_ = &Type{
		"const QIncompatibleFlag&",
		classes["QIncompatibleFlag"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[552] = type_
	type_ = &Type{
		"const QInternal&",
		classes["QInternal"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[553] = type_
	type_ = &Type{
		"const QLatin1Char&",
		classes["QLatin1Char"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[554] = type_
	type_ = &Type{
		"const QLatin1String&",
		classes["QLatin1String"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[555] = type_
	type_ = &Type{
		"const QLibraryInfo&",
		classes["QLibraryInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[556] = type_
	type_ = &Type{
		"const QLine&",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[557] = type_
	type_ = &Type{
		"const QLineF&",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[558] = type_
	type_ = &Type{
		"const QList<QAbstractState*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[559] = type_
	type_ = &Type{
		"const QList<QModelIndex>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[560] = type_
	type_ = &Type{
		"const QList<QObject*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[561] = type_
	type_ = &Type{
		"const QList<QPair<QByteArray,QByteArray> >&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[562] = type_
	type_ = &Type{
		"const QList<QPair<QString,QString> >&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[563] = type_
	type_ = &Type{
		"const QList<QUrl>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[564] = type_
	type_ = &Type{
		"const QList<QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[565] = type_
	type_ = &Type{
		"const QLocale&",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[566] = type_
	type_ = &Type{
		"const QMap<QString,QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[567] = type_
	type_ = &Type{
		"const QMap<int,QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[568] = type_
	type_ = &Type{
		"const QMargins&",
		classes["QMargins"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[569] = type_
	type_ = &Type{
		"const QMetaClassInfo&",
		classes["QMetaClassInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[570] = type_
	type_ = &Type{
		"const QMetaEnum&",
		classes["QMetaEnum"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[571] = type_
	type_ = &Type{
		"const QMetaMethod&",
		classes["QMetaMethod"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[572] = type_
	type_ = &Type{
		"const QMetaObject&",
		classes["QMetaObject"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[573] = type_
	type_ = &Type{
		"const QMetaObject*",
		classes["QMetaObject"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[574] = type_
	type_ = &Type{
		"const QMetaProperty&",
		classes["QMetaProperty"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[575] = type_
	type_ = &Type{
		"const QMetaType&",
		classes["QMetaType"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[576] = type_
	type_ = &Type{
		"const QMimeData*",
		classes["QMimeData"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[577] = type_
	type_ = &Type{
		"const QModelIndex&",
		classes["QModelIndex"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[578] = type_
	type_ = &Type{
		"const QNoDebug&",
		classes["QNoDebug"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[579] = type_
	type_ = &Type{
		"const QObject*",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[580] = type_
	type_ = &Type{
		"const QObjectUserData&",
		classes["QObjectUserData"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[581] = type_
	type_ = &Type{
		"const QPersistentModelIndex&",
		classes["QPersistentModelIndex"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[582] = type_
	type_ = &Type{
		"const QPoint",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[583] = type_
	type_ = &Type{
		"const QPoint&",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[584] = type_
	type_ = &Type{
		"const QPointF",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[585] = type_
	type_ = &Type{
		"const QPointF&",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[586] = type_
	type_ = &Type{
		"const QProcessEnvironment&",
		classes["QProcessEnvironment"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[587] = type_
	type_ = &Type{
		"const QRect&",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[588] = type_
	type_ = &Type{
		"const QRectF&",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[589] = type_
	type_ = &Type{
		"const QRegExp&",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[590] = type_
	type_ = &Type{
		"const QRegExp*",
		classes["QRegExp"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[591] = type_
	type_ = &Type{
		"const QRunnable&",
		classes["QRunnable"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[592] = type_
	type_ = &Type{
		"const QSharedData&",
		classes["QSharedData"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[593] = type_
	type_ = &Type{
		"const QSize",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[594] = type_
	type_ = &Type{
		"const QSize&",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[595] = type_
	type_ = &Type{
		"const QSizeF",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[596] = type_
	type_ = &Type{
		"const QSizeF&",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[597] = type_
	type_ = &Type{
		"const QString",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[598] = type_
	type_ = &Type{
		"const QString&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[599] = type_
	type_ = &Type{
		"const QString*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[600] = type_
	type_ = &Type{
		"const QStringList&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[601] = type_
	type_ = &Type{
		"const QStringList*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[602] = type_
	type_ = &Type{
		"const QStringMatcher&",
		classes["QStringMatcher"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[603] = type_
	type_ = &Type{
		"const QStringRef&",
		classes["QStringRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[604] = type_
	type_ = &Type{
		"const QSysInfo&",
		classes["QSysInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[605] = type_
	type_ = &Type{
		"const QSystemLocale&",
		classes["QSystemLocale"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[606] = type_
	type_ = &Type{
		"const QTextBoundaryFinder&",
		classes["QTextBoundaryFinder"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[607] = type_
	type_ = &Type{
		"const QTextCodec*",
		classes["QTextCodec"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[608] = type_
	type_ = &Type{
		"const QTextCodecFactoryInterface&",
		classes["QTextCodecFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[609] = type_
	type_ = &Type{
		"const QTime&",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[610] = type_
	type_ = &Type{
		"const QTimerEvent&",
		classes["QTimerEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[611] = type_
	type_ = &Type{
		"const QUrl&",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[612] = type_
	type_ = &Type{
		"const QUuid&",
		classes["QUuid"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[613] = type_
	type_ = &Type{
		"const QVariant&",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[614] = type_
	type_ = &Type{
		"const QVariant::Handler*",
		classes["QVariant::Handler"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[615] = type_
	type_ = &Type{
		"const QVariant::Private&",
		classes["QVariant::Private"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[616] = type_
	type_ = &Type{
		"const QVariant::Type",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		true,
	}
	types[617] = type_
	type_ = &Type{
		"const QVariantComparisonHelper&",
		classes["QVariantComparisonHelper"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[618] = type_
	type_ = &Type{
		"const QVector<QPair<double,QVariant> >&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[619] = type_
	type_ = &Type{
		"const QVector<QXmlStreamNamespaceDeclaration>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[620] = type_
	type_ = &Type{
		"const QXmlStreamAttribute&",
		classes["QXmlStreamAttribute"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[621] = type_
	type_ = &Type{
		"const QXmlStreamAttributes&",
		classes["QXmlStreamAttributes"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[622] = type_
	type_ = &Type{
		"const QXmlStreamEntityDeclaration&",
		classes["QXmlStreamEntityDeclaration"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[623] = type_
	type_ = &Type{
		"const QXmlStreamEntityResolver&",
		classes["QXmlStreamEntityResolver"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[624] = type_
	type_ = &Type{
		"const QXmlStreamNamespaceDeclaration&",
		classes["QXmlStreamNamespaceDeclaration"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[625] = type_
	type_ = &Type{
		"const QXmlStreamNotationDeclaration&",
		classes["QXmlStreamNotationDeclaration"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[626] = type_
	type_ = &Type{
		"const QXmlStreamReader&",
		classes["QXmlStreamReader"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[627] = type_
	type_ = &Type{
		"const QXmlStreamStringRef&",
		classes["QXmlStreamStringRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[628] = type_
	type_ = &Type{
		"const char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[629] = type_
	type_ = &Type{
		"const unsigned char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[630] = type_
	type_ = &Type{
		"const void*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[631] = type_
	type_ = &Type{
		"double",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[632] = type_
	type_ = &Type{
		"double&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[633] = type_
	type_ = &Type{
		"double(*)(double)",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[634] = type_
	type_ = &Type{
		"double*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[635] = type_
	type_ = &Type{
		"float",
		nil,
		T_FLOAT,
		KIND_STACK,
		false,
	}
	types[636] = type_
	type_ = &Type{
		"float&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[637] = type_
	type_ = &Type{
		"int",
		nil,
		T_INT,
		KIND_STACK,
		false,
	}
	types[638] = type_
	type_ = &Type{
		"int&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[639] = type_
	type_ = &Type{
		"int(*)(const void*,const void*)",
		nil,
		T_INT,
		KIND_STACK,
		false,
	}
	types[640] = type_
	type_ = &Type{
		"int*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[641] = type_
	type_ = &Type{
		"long",
		nil,
		T_LONG,
		KIND_STACK,
		false,
	}
	types[642] = type_
	type_ = &Type{
		"long double",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[643] = type_
	type_ = &Type{
		"long long",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[644] = type_
	type_ = &Type{
		"long long&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[645] = type_
	type_ = &Type{
		"long long*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[646] = type_
	type_ = &Type{
		"long*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[647] = type_
	type_ = &Type{
		"qint64",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[648] = type_
	type_ = &Type{
		"qreal",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[649] = type_
	type_ = &Type{
		"short",
		nil,
		T_SHORT,
		KIND_STACK,
		false,
	}
	types[650] = type_
	type_ = &Type{
		"short&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[651] = type_
	type_ = &Type{
		"signed char",
		nil,
		T_CHAR,
		KIND_STACK,
		false,
	}
	types[652] = type_
	type_ = &Type{
		"signed char&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[653] = type_
	type_ = &Type{
		"signed int",
		nil,
		T_INT,
		KIND_STACK,
		false,
	}
	types[654] = type_
	type_ = &Type{
		"signed int&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[655] = type_
	type_ = &Type{
		"signed long",
		nil,
		T_LONG,
		KIND_STACK,
		false,
	}
	types[656] = type_
	type_ = &Type{
		"signed long&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[657] = type_
	type_ = &Type{
		"signed short",
		nil,
		T_SHORT,
		KIND_STACK,
		false,
	}
	types[658] = type_
	type_ = &Type{
		"signed short&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[659] = type_
	type_ = &Type{
		"size_t",
		nil,
		T_ULONG,
		KIND_STACK,
		false,
	}
	types[660] = type_
	type_ = &Type{
		"unsigned char",
		nil,
		T_UCHAR,
		KIND_STACK,
		false,
	}
	types[661] = type_
	type_ = &Type{
		"unsigned char&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[662] = type_
	type_ = &Type{
		"unsigned char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[663] = type_
	type_ = &Type{
		"unsigned int",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[664] = type_
	type_ = &Type{
		"unsigned int&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[665] = type_
	type_ = &Type{
		"unsigned long",
		nil,
		T_ULONG,
		KIND_STACK,
		false,
	}
	types[666] = type_
	type_ = &Type{
		"unsigned long long",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[667] = type_
	type_ = &Type{
		"unsigned long long&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[668] = type_
	type_ = &Type{
		"unsigned long&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[669] = type_
	type_ = &Type{
		"unsigned short",
		nil,
		T_USHORT,
		KIND_STACK,
		false,
	}
	types[670] = type_
	type_ = &Type{
		"unsigned short&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[671] = type_
	type_ = &Type{
		"va_list",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[672] = type_
	type_ = &Type{
		"void(*)()",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[673] = type_
	type_ = &Type{
		"void(*)(QDataStream&,const void*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[674] = type_
	type_ = &Type{
		"void(*)(QDataStream&,void*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[675] = type_
	type_ = &Type{
		"void(*)(QtMsgType,const char*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[676] = type_
	type_ = &Type{
		"void(*)(void*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[677] = type_
	type_ = &Type{
		"void*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[678] = type_
	type_ = &Type{
		"void*(*)(const void*)",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[679] = type_
	type_ = &Type{
		"void**",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[680] = type_
	type_ = &Type{
		"volatile const void*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[681] = type_
	klass = &Class{
		"QAbstractButton",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractButton"] = klass
	klass = &Class{
		"QAbstractGraphicsShapeItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractGraphicsShapeItem"] = klass
	klass = &Class{
		"QAbstractItemDelegate",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractItemDelegate"] = klass
	klass = &Class{
		"QAbstractItemModel",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QAbstractItemModel"] = klass
	klass = &Class{
		"QAbstractItemView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractItemView"] = klass
	klass = &Class{
		"QAbstractListModel",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QAbstractListModel"] = klass
	klass = &Class{
		"QAbstractPageSetupDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractPageSetupDialog"] = klass
	klass = &Class{
		"QAbstractPrintDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractPrintDialog"] = klass
	klass = &Class{
		"QAbstractProxyModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractProxyModel"] = klass
	klass = &Class{
		"QAbstractScrollArea",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractScrollArea"] = klass
	klass = &Class{
		"QAbstractSlider",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractSlider"] = klass
	klass = &Class{
		"QAbstractSpinBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractSpinBox"] = klass
	klass = &Class{
		"QAbstractTextDocumentLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAbstractTextDocumentLayout"] = klass
	klass = &Class{
		"QAbstractTextDocumentLayout::PaintContext",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractTextDocumentLayout::PaintContext"] = klass
	klass = &Class{
		"QAbstractTextDocumentLayout::Selection",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAbstractTextDocumentLayout::Selection"] = klass
	klass = &Class{
		"QAbstractUndoItem",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAbstractUndoItem"] = klass
	klass = &Class{
		"QAccessible",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAccessible"] = klass
	klass = &Class{
		"QAccessible2",
		false,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		true,
		false,
	}
	classes["QAccessible2"] = klass
	klass = &Class{
		"QAccessible2::TableModelChange",
		false,
		make([]*Class, 0),
		20,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QAccessible2::TableModelChange"] = klass
	klass = &Class{
		"QAccessible2Interface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessible2Interface"] = klass
	klass = &Class{
		"QAccessibleActionInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleActionInterface"] = klass
	klass = &Class{
		"QAccessibleApplication",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleApplication"] = klass
	klass = &Class{
		"QAccessibleBridge",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleBridge"] = klass
	klass = &Class{
		"QAccessibleBridgeFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleBridgeFactoryInterface"] = klass
	klass = &Class{
		"QAccessibleBridgePlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleBridgePlugin"] = klass
	klass = &Class{
		"QAccessibleEditableTextInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleEditableTextInterface"] = klass
	klass = &Class{
		"QAccessibleEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleEvent"] = klass
	klass = &Class{
		"QAccessibleFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleFactoryInterface"] = klass
	klass = &Class{
		"QAccessibleImageInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleImageInterface"] = klass
	klass = &Class{
		"QAccessibleInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleInterface"] = klass
	klass = &Class{
		"QAccessibleInterfaceEx",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleInterfaceEx"] = klass
	klass = &Class{
		"QAccessibleObject",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleObject"] = klass
	klass = &Class{
		"QAccessibleObjectEx",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleObjectEx"] = klass
	klass = &Class{
		"QAccessiblePlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessiblePlugin"] = klass
	klass = &Class{
		"QAccessibleSimpleEditableTextInterface",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleSimpleEditableTextInterface"] = klass
	klass = &Class{
		"QAccessibleTable2CellInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleTable2CellInterface"] = klass
	klass = &Class{
		"QAccessibleTable2Interface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleTable2Interface"] = klass
	klass = &Class{
		"QAccessibleTableInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleTableInterface"] = klass
	klass = &Class{
		"QAccessibleTextInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleTextInterface"] = klass
	klass = &Class{
		"QAccessibleValueInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QAccessibleValueInterface"] = klass
	klass = &Class{
		"QAccessibleWidget",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleWidget"] = klass
	klass = &Class{
		"QAccessibleWidgetEx",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAccessibleWidgetEx"] = klass
	klass = &Class{
		"QAction",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QAction"] = klass
	klass = &Class{
		"QActionEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QActionEvent"] = klass
	klass = &Class{
		"QActionGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QActionGroup"] = klass
	klass = &Class{
		"QApplication",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QApplication"] = klass
	klass = &Class{
		"QBitArray",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QBitArray"] = klass
	klass = &Class{
		"QBitmap",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QBitmap"] = klass
	klass = &Class{
		"QBool",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QBool"] = klass
	klass = &Class{
		"QBoxLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QBoxLayout"] = klass
	klass = &Class{
		"QBrush",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QBrush"] = klass
	klass = &Class{
		"QButtonGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QButtonGroup"] = klass
	klass = &Class{
		"QByteArray",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QByteArray"] = klass
	klass = &Class{
		"QCalendarWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCalendarWidget"] = klass
	klass = &Class{
		"QChar",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QChar"] = klass
	klass = &Class{
		"QCheckBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCheckBox"] = klass
	klass = &Class{
		"QChildEvent",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QChildEvent"] = klass
	klass = &Class{
		"QClipboard",
		false,
		make([]*Class, 0),
		16,
		false,
		false,
		true,
		false,
		false,
	}
	classes["QClipboard"] = klass
	klass = &Class{
		"QClipboardEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QClipboardEvent"] = klass
	klass = &Class{
		"QCloseEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QCloseEvent"] = klass
	klass = &Class{
		"QColor",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QColor"] = klass
	klass = &Class{
		"QColorDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QColorDialog"] = klass
	klass = &Class{
		"QColormap",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QColormap"] = klass
	klass = &Class{
		"QColumnView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QColumnView"] = klass
	klass = &Class{
		"QComboBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QComboBox"] = klass
	klass = &Class{
		"QCommandLinkButton",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCommandLinkButton"] = klass
	klass = &Class{
		"QCommonStyle",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCommonStyle"] = klass
	klass = &Class{
		"QCompleter",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QCompleter"] = klass
	klass = &Class{
		"QConicalGradient",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QConicalGradient"] = klass
	klass = &Class{
		"QContextMenuEvent",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QContextMenuEvent"] = klass
	klass = &Class{
		"QCoreApplication",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QCoreApplication"] = klass
	klass = &Class{
		"QCursor",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QCursor"] = klass
	klass = &Class{
		"QDataStream",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QDataStream"] = klass
	klass = &Class{
		"QDataWidgetMapper",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDataWidgetMapper"] = klass
	klass = &Class{
		"QDate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QDate"] = klass
	klass = &Class{
		"QDateEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDateEdit"] = klass
	klass = &Class{
		"QDateTime",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QDateTime"] = klass
	klass = &Class{
		"QDateTimeEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDateTimeEdit"] = klass
	klass = &Class{
		"QDebug",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QDebug"] = klass
	klass = &Class{
		"QDesktopServices",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QDesktopServices"] = klass
	klass = &Class{
		"QDesktopWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDesktopWidget"] = klass
	klass = &Class{
		"QDial",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDial"] = klass
	klass = &Class{
		"QDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDialog"] = klass
	klass = &Class{
		"QDialogButtonBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDialogButtonBox"] = klass
	klass = &Class{
		"QDir",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QDir"] = klass
	klass = &Class{
		"QDirModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDirModel"] = klass
	klass = &Class{
		"QDockWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDockWidget"] = klass
	klass = &Class{
		"QDoubleSpinBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDoubleSpinBox"] = klass
	klass = &Class{
		"QDoubleValidator",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDoubleValidator"] = klass
	klass = &Class{
		"QDrag",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QDrag"] = klass
	klass = &Class{
		"QDragEnterEvent",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDragEnterEvent"] = klass
	klass = &Class{
		"QDragLeaveEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDragLeaveEvent"] = klass
	klass = &Class{
		"QDragMoveEvent",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDragMoveEvent"] = klass
	klass = &Class{
		"QDragResponseEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDragResponseEvent"] = klass
	klass = &Class{
		"QDropEvent",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QDropEvent"] = klass
	klass = &Class{
		"QEasingCurve",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QEasingCurve"] = klass
	klass = &Class{
		"QErrorMessage",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QErrorMessage"] = klass
	klass = &Class{
		"QEvent",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QEvent"] = klass
	klass = &Class{
		"QEventPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QEventPrivate"] = klass
	klass = &Class{
		"QEventTransition",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QEventTransition"] = klass
	klass = &Class{
		"QFactoryInterface",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QFactoryInterface"] = klass
	klass = &Class{
		"QFile",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QFile"] = klass
	klass = &Class{
		"QFileDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFileDialog"] = klass
	klass = &Class{
		"QFileDialogArgs",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QFileDialogArgs"] = klass
	klass = &Class{
		"QFileIconProvider",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFileIconProvider"] = klass
	klass = &Class{
		"QFileInfo",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QFileInfo"] = klass
	klass = &Class{
		"QFileOpenEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QFileOpenEvent"] = klass
	klass = &Class{
		"QFileSystemModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFileSystemModel"] = klass
	klass = &Class{
		"QFocusEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QFocusEvent"] = klass
	klass = &Class{
		"QFocusFrame",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFocusFrame"] = klass
	klass = &Class{
		"QFont",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFont"] = klass
	klass = &Class{
		"QFontComboBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFontComboBox"] = klass
	klass = &Class{
		"QFontDatabase",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFontDatabase"] = klass
	klass = &Class{
		"QFontDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFontDialog"] = klass
	klass = &Class{
		"QFontInfo",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFontInfo"] = klass
	klass = &Class{
		"QFontMetrics",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFontMetrics"] = klass
	klass = &Class{
		"QFontMetricsF",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QFontMetricsF"] = klass
	klass = &Class{
		"QFormLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFormLayout"] = klass
	klass = &Class{
		"QFrame",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QFrame"] = klass
	klass = &Class{
		"QGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGesture"] = klass
	klass = &Class{
		"QGestureEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QGestureEvent"] = klass
	klass = &Class{
		"QGestureRecognizer",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QGestureRecognizer"] = klass
	klass = &Class{
		"QGlobalSpace",
		false,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		true,
		false,
	}
	classes["QGlobalSpace"] = klass
	klass = &Class{
		"QGlyphRun",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QGlyphRun"] = klass
	klass = &Class{
		"QGradient",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QGradient"] = klass
	klass = &Class{
		"QGraphicsAnchor",
		false,
		make([]*Class, 0),
		16,
		false,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsAnchor"] = klass
	klass = &Class{
		"QGraphicsAnchorLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsAnchorLayout"] = klass
	klass = &Class{
		"QGraphicsBlurEffect",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsBlurEffect"] = klass
	klass = &Class{
		"QGraphicsColorizeEffect",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsColorizeEffect"] = klass
	klass = &Class{
		"QGraphicsDropShadowEffect",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsDropShadowEffect"] = klass
	klass = &Class{
		"QGraphicsEffect",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsEffect"] = klass
	klass = &Class{
		"QGraphicsEffectSource",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QGraphicsEffectSource"] = klass
	klass = &Class{
		"QGraphicsEllipseItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsEllipseItem"] = klass
	klass = &Class{
		"QGraphicsGridLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsGridLayout"] = klass
	klass = &Class{
		"QGraphicsItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsItem"] = klass
	klass = &Class{
		"QGraphicsItemAnimation",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsItemAnimation"] = klass
	klass = &Class{
		"QGraphicsItemGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsItemGroup"] = klass
	klass = &Class{
		"QGraphicsLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsLayout"] = klass
	klass = &Class{
		"QGraphicsLayoutItem",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QGraphicsLayoutItem"] = klass
	klass = &Class{
		"QGraphicsLineItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsLineItem"] = klass
	klass = &Class{
		"QGraphicsLinearLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsLinearLayout"] = klass
	klass = &Class{
		"QGraphicsObject",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsObject"] = klass
	klass = &Class{
		"QGraphicsOpacityEffect",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsOpacityEffect"] = klass
	klass = &Class{
		"QGraphicsPathItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsPathItem"] = klass
	klass = &Class{
		"QGraphicsPixmapItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsPixmapItem"] = klass
	klass = &Class{
		"QGraphicsPolygonItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsPolygonItem"] = klass
	klass = &Class{
		"QGraphicsProxyWidget",
		false,
		make([]*Class, 0),
		48,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsProxyWidget"] = klass
	klass = &Class{
		"QGraphicsRectItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsRectItem"] = klass
	klass = &Class{
		"QGraphicsRotation",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsRotation"] = klass
	klass = &Class{
		"QGraphicsScale",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsScale"] = klass
	klass = &Class{
		"QGraphicsScene",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsScene"] = klass
	klass = &Class{
		"QGraphicsSceneContextMenuEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneContextMenuEvent"] = klass
	klass = &Class{
		"QGraphicsSceneDragDropEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneDragDropEvent"] = klass
	klass = &Class{
		"QGraphicsSceneEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneEvent"] = klass
	klass = &Class{
		"QGraphicsSceneEventPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QGraphicsSceneEventPrivate"] = klass
	klass = &Class{
		"QGraphicsSceneHelpEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneHelpEvent"] = klass
	klass = &Class{
		"QGraphicsSceneHoverEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneHoverEvent"] = klass
	klass = &Class{
		"QGraphicsSceneMouseEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneMouseEvent"] = klass
	klass = &Class{
		"QGraphicsSceneMoveEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneMoveEvent"] = klass
	klass = &Class{
		"QGraphicsSceneResizeEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneResizeEvent"] = klass
	klass = &Class{
		"QGraphicsSceneWheelEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSceneWheelEvent"] = klass
	klass = &Class{
		"QGraphicsSimpleTextItem",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsSimpleTextItem"] = klass
	klass = &Class{
		"QGraphicsTextItem",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsTextItem"] = klass
	klass = &Class{
		"QGraphicsTransform",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsTransform"] = klass
	klass = &Class{
		"QGraphicsView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsView"] = klass
	klass = &Class{
		"QGraphicsWidget",
		false,
		make([]*Class, 0),
		48,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGraphicsWidget"] = klass
	klass = &Class{
		"QGridLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGridLayout"] = klass
	klass = &Class{
		"QGroupBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QGroupBox"] = klass
	klass = &Class{
		"QHBoxLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QHBoxLayout"] = klass
	klass = &Class{
		"QHashDummyValue",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QHashDummyValue"] = klass
	klass = &Class{
		"QHeaderView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QHeaderView"] = klass
	klass = &Class{
		"QHelpEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QHelpEvent"] = klass
	klass = &Class{
		"QHideEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QHideEvent"] = klass
	klass = &Class{
		"QHoverEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QHoverEvent"] = klass
	klass = &Class{
		"QIODevice",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QIODevice"] = klass
	klass = &Class{
		"QIcon",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QIcon"] = klass
	klass = &Class{
		"QIconDragEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QIconDragEvent"] = klass
	klass = &Class{
		"QIconEngine",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QIconEngine"] = klass
	klass = &Class{
		"QIconEngineFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QIconEngineFactoryInterface"] = klass
	klass = &Class{
		"QIconEngineFactoryInterfaceV2",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QIconEngineFactoryInterfaceV2"] = klass
	klass = &Class{
		"QIconEnginePlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QIconEnginePlugin"] = klass
	klass = &Class{
		"QIconEnginePluginV2",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QIconEnginePluginV2"] = klass
	klass = &Class{
		"QIconEngineV2",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QIconEngineV2"] = klass
	klass = &Class{
		"QIconEngineV2::AvailableSizesArgument",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QIconEngineV2::AvailableSizesArgument"] = klass
	klass = &Class{
		"QIconPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QIconPrivate"] = klass
	klass = &Class{
		"QImage",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QImage"] = klass
	klass = &Class{
		"QImageData",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QImageData"] = klass
	klass = &Class{
		"QImageIOHandler",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QImageIOHandler"] = klass
	klass = &Class{
		"QImageIOHandlerFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QImageIOHandlerFactoryInterface"] = klass
	klass = &Class{
		"QImageIOPlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QImageIOPlugin"] = klass
	klass = &Class{
		"QImageReader",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QImageReader"] = klass
	klass = &Class{
		"QImageTextKeyLang",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QImageTextKeyLang"] = klass
	klass = &Class{
		"QImageWriter",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QImageWriter"] = klass
	klass = &Class{
		"QIncompatibleFlag",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QIncompatibleFlag"] = klass
	klass = &Class{
		"QInputContext",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QInputContext"] = klass
	klass = &Class{
		"QInputContextFactory",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QInputContextFactory"] = klass
	klass = &Class{
		"QInputContextFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QInputContextFactoryInterface"] = klass
	klass = &Class{
		"QInputContextPlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QInputContextPlugin"] = klass
	klass = &Class{
		"QInputDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QInputDialog"] = klass
	klass = &Class{
		"QInputEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QInputEvent"] = klass
	klass = &Class{
		"QInputMethodEvent",
		false,
		make([]*Class, 0),
		56,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QInputMethodEvent"] = klass
	klass = &Class{
		"QInputMethodEvent::Attribute",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QInputMethodEvent::Attribute"] = klass
	klass = &Class{
		"QIntValidator",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QIntValidator"] = klass
	klass = &Class{
		"QItemDelegate",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QItemDelegate"] = klass
	klass = &Class{
		"QItemEditorCreatorBase",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QItemEditorCreatorBase"] = klass
	klass = &Class{
		"QItemEditorFactory",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QItemEditorFactory"] = klass
	klass = &Class{
		"QItemSelection",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QItemSelection"] = klass
	klass = &Class{
		"QItemSelectionModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QItemSelectionModel"] = klass
	klass = &Class{
		"QItemSelectionRange",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QItemSelectionRange"] = klass
	klass = &Class{
		"QKeyEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QKeyEvent"] = klass
	klass = &Class{
		"QKeyEventTransition",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QKeyEventTransition"] = klass
	klass = &Class{
		"QKeySequence",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QKeySequence"] = klass
	klass = &Class{
		"QKeySequencePrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QKeySequencePrivate"] = klass
	klass = &Class{
		"QLCDNumber",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QLCDNumber"] = klass
	klass = &Class{
		"QLabel",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QLabel"] = klass
	klass = &Class{
		"QLatin1String",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QLatin1String"] = klass
	klass = &Class{
		"QLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QLayout"] = klass
	klass = &Class{
		"QLayoutItem",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QLayoutItem"] = klass
	klass = &Class{
		"QLine",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QLine"] = klass
	klass = &Class{
		"QLineEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QLineEdit"] = klass
	klass = &Class{
		"QLineF",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QLineF"] = klass
	klass = &Class{
		"QLinearGradient",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QLinearGradient"] = klass
	klass = &Class{
		"QListView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QListView"] = klass
	klass = &Class{
		"QListWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QListWidget"] = klass
	klass = &Class{
		"QListWidgetItem",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QListWidgetItem"] = klass
	klass = &Class{
		"QLocale",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QLocale"] = klass
	klass = &Class{
		"QMainWindow",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMainWindow"] = klass
	klass = &Class{
		"QMargins",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QMargins"] = klass
	klass = &Class{
		"QMatrix",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMatrix"] = klass
	klass = &Class{
		"QMatrix4x4",
		false,
		make([]*Class, 0),
		136,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QMatrix4x4"] = klass
	klass = &Class{
		"QMdiArea",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMdiArea"] = klass
	klass = &Class{
		"QMdiSubWindow",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMdiSubWindow"] = klass
	klass = &Class{
		"QMenu",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMenu"] = klass
	klass = &Class{
		"QMenuBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMenuBar"] = klass
	klass = &Class{
		"QMessageBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMessageBox"] = klass
	klass = &Class{
		"QMetaObject",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QMetaObject"] = klass
	klass = &Class{
		"QMimeData",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QMimeData"] = klass
	klass = &Class{
		"QMimeSource",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QMimeSource"] = klass
	klass = &Class{
		"QModelIndex",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QModelIndex"] = klass
	klass = &Class{
		"QMouseEvent",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QMouseEvent"] = klass
	klass = &Class{
		"QMouseEventTransition",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMouseEventTransition"] = klass
	klass = &Class{
		"QMoveEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QMoveEvent"] = klass
	klass = &Class{
		"QMovie",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QMovie"] = klass
	klass = &Class{
		"QObject",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QObject"] = klass
	klass = &Class{
		"QPageSetupDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPageSetupDialog"] = klass
	klass = &Class{
		"QPaintDevice",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPaintDevice"] = klass
	klass = &Class{
		"QPaintEngine",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPaintEngine"] = klass
	klass = &Class{
		"QPaintEngineState",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPaintEngineState"] = klass
	klass = &Class{
		"QPaintEvent",
		false,
		make([]*Class, 0),
		56,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QPaintEvent"] = klass
	klass = &Class{
		"QPainter",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QPainter"] = klass
	klass = &Class{
		"QPainter::PixmapFragment",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPainter::PixmapFragment"] = klass
	klass = &Class{
		"QPainterPath",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPainterPath"] = klass
	klass = &Class{
		"QPainterPath::Element",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPainterPath::Element"] = klass
	klass = &Class{
		"QPainterPathStroker",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QPainterPathStroker"] = klass
	klass = &Class{
		"QPalette",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPalette"] = klass
	klass = &Class{
		"QPanGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPanGesture"] = klass
	klass = &Class{
		"QPen",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPen"] = klass
	klass = &Class{
		"QPenPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPenPrivate"] = klass
	klass = &Class{
		"QPersistentModelIndex",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPersistentModelIndex"] = klass
	klass = &Class{
		"QPicture",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPicture"] = klass
	klass = &Class{
		"QPictureFormatInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QPictureFormatInterface"] = klass
	klass = &Class{
		"QPictureFormatPlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPictureFormatPlugin"] = klass
	klass = &Class{
		"QPictureIO",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QPictureIO"] = klass
	klass = &Class{
		"QPinchGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPinchGesture"] = klass
	klass = &Class{
		"QPixmap",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPixmap"] = klass
	klass = &Class{
		"QPixmapCache",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPixmapCache"] = klass
	klass = &Class{
		"QPixmapCache::Key",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPixmapCache::Key"] = klass
	klass = &Class{
		"QPlainTextDocumentLayout",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPlainTextDocumentLayout"] = klass
	klass = &Class{
		"QPlainTextEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPlainTextEdit"] = klass
	klass = &Class{
		"QPoint",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPoint"] = klass
	klass = &Class{
		"QPointF",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPointF"] = klass
	klass = &Class{
		"QPolygon",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPolygon"] = klass
	klass = &Class{
		"QPolygonF",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPolygonF"] = klass
	klass = &Class{
		"QPostEventList",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QPostEventList"] = klass
	klass = &Class{
		"QPrintDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPrintDialog"] = klass
	klass = &Class{
		"QPrintEngine",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QPrintEngine"] = klass
	klass = &Class{
		"QPrintPreviewDialog",
		false,
		make([]*Class, 0),
		48,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPrintPreviewDialog"] = klass
	klass = &Class{
		"QPrintPreviewWidget",
		false,
		make([]*Class, 0),
		48,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPrintPreviewWidget"] = klass
	klass = &Class{
		"QPrinter",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPrinter"] = klass
	klass = &Class{
		"QPrinterInfo",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QPrinterInfo"] = klass
	klass = &Class{
		"QProgressBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QProgressBar"] = klass
	klass = &Class{
		"QProgressDialog",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QProgressDialog"] = klass
	klass = &Class{
		"QProxyModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QProxyModel"] = klass
	klass = &Class{
		"QProxyStyle",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QProxyStyle"] = klass
	klass = &Class{
		"QPushButton",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QPushButton"] = klass
	klass = &Class{
		"QQuaternion",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QQuaternion"] = klass
	klass = &Class{
		"QRadialGradient",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRadialGradient"] = klass
	klass = &Class{
		"QRadioButton",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QRadioButton"] = klass
	klass = &Class{
		"QRawFont",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRawFont"] = klass
	klass = &Class{
		"QRect",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QRect"] = klass
	klass = &Class{
		"QRectF",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QRectF"] = klass
	klass = &Class{
		"QRegExp",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QRegExp"] = klass
	klass = &Class{
		"QRegExpValidator",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QRegExpValidator"] = klass
	klass = &Class{
		"QRegion",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QRegion"] = klass
	klass = &Class{
		"QResizeEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QResizeEvent"] = klass
	klass = &Class{
		"QRubberBand",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QRubberBand"] = klass
	klass = &Class{
		"QScrollArea",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QScrollArea"] = klass
	klass = &Class{
		"QScrollBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QScrollBar"] = klass
	klass = &Class{
		"QSessionManager",
		false,
		make([]*Class, 0),
		16,
		false,
		false,
		true,
		false,
		false,
	}
	classes["QSessionManager"] = klass
	klass = &Class{
		"QShortcut",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QShortcut"] = klass
	klass = &Class{
		"QShortcutEvent",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QShortcutEvent"] = klass
	klass = &Class{
		"QShowEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QShowEvent"] = klass
	klass = &Class{
		"QSize",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QSize"] = klass
	klass = &Class{
		"QSizeF",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QSizeF"] = klass
	klass = &Class{
		"QSizeGrip",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSizeGrip"] = klass
	klass = &Class{
		"QSizePolicy",
		false,
		make([]*Class, 0),
		4,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QSizePolicy"] = klass
	klass = &Class{
		"QSlider",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSlider"] = klass
	klass = &Class{
		"QSortFilterProxyModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSortFilterProxyModel"] = klass
	klass = &Class{
		"QSound",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSound"] = klass
	klass = &Class{
		"QSpacerItem",
		false,
		make([]*Class, 0),
		40,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QSpacerItem"] = klass
	klass = &Class{
		"QSpinBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSpinBox"] = klass
	klass = &Class{
		"QSplashScreen",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSplashScreen"] = klass
	klass = &Class{
		"QSplitter",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSplitter"] = klass
	klass = &Class{
		"QSplitterHandle",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSplitterHandle"] = klass
	klass = &Class{
		"QStackedLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStackedLayout"] = klass
	klass = &Class{
		"QStackedWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStackedWidget"] = klass
	klass = &Class{
		"QStandardItem",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QStandardItem"] = klass
	klass = &Class{
		"QStandardItemModel",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStandardItemModel"] = klass
	klass = &Class{
		"QState",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QState"] = klass
	klass = &Class{
		"QStaticText",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStaticText"] = klass
	klass = &Class{
		"QStatusBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStatusBar"] = klass
	klass = &Class{
		"QStatusTipEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QStatusTipEvent"] = klass
	klass = &Class{
		"QString::Null",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QString::Null"] = klass
	klass = &Class{
		"QStringListModel",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStringListModel"] = klass
	klass = &Class{
		"QStringRef",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QStringRef"] = klass
	klass = &Class{
		"QStyle",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStyle"] = klass
	klass = &Class{
		"QStyleFactory",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleFactory"] = klass
	klass = &Class{
		"QStyleFactoryInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QStyleFactoryInterface"] = klass
	klass = &Class{
		"QStyleHintReturn",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleHintReturn"] = klass
	klass = &Class{
		"QStyleHintReturnMask",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleHintReturnMask"] = klass
	klass = &Class{
		"QStyleHintReturnVariant",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleHintReturnVariant"] = klass
	klass = &Class{
		"QStyleOption",
		false,
		make([]*Class, 0),
		56,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOption"] = klass
	klass = &Class{
		"QStyleOptionButton",
		false,
		make([]*Class, 0),
		88,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionButton"] = klass
	klass = &Class{
		"QStyleOptionComboBox",
		false,
		make([]*Class, 0),
		112,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionComboBox"] = klass
	klass = &Class{
		"QStyleOptionComplex",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionComplex"] = klass
	klass = &Class{
		"QStyleOptionDockWidget",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionDockWidget"] = klass
	klass = &Class{
		"QStyleOptionDockWidgetV2",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionDockWidgetV2"] = klass
	klass = &Class{
		"QStyleOptionFocusRect",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionFocusRect"] = klass
	klass = &Class{
		"QStyleOptionFrame",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionFrame"] = klass
	klass = &Class{
		"QStyleOptionFrameV2",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionFrameV2"] = klass
	klass = &Class{
		"QStyleOptionFrameV3",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionFrameV3"] = klass
	klass = &Class{
		"QStyleOptionGraphicsItem",
		false,
		make([]*Class, 0),
		144,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionGraphicsItem"] = klass
	klass = &Class{
		"QStyleOptionGroupBox",
		false,
		make([]*Class, 0),
		112,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionGroupBox"] = klass
	klass = &Class{
		"QStyleOptionHeader",
		false,
		make([]*Class, 0),
		112,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionHeader"] = klass
	klass = &Class{
		"QStyleOptionMenuItem",
		false,
		make([]*Class, 0),
		128,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionMenuItem"] = klass
	klass = &Class{
		"QStyleOptionProgressBar",
		false,
		make([]*Class, 0),
		88,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionProgressBar"] = klass
	klass = &Class{
		"QStyleOptionProgressBarV2",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionProgressBarV2"] = klass
	klass = &Class{
		"QStyleOptionRubberBand",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionRubberBand"] = klass
	klass = &Class{
		"QStyleOptionSizeGrip",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionSizeGrip"] = klass
	klass = &Class{
		"QStyleOptionSlider",
		false,
		make([]*Class, 0),
		120,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionSlider"] = klass
	klass = &Class{
		"QStyleOptionSpinBox",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionSpinBox"] = klass
	klass = &Class{
		"QStyleOptionTab",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTab"] = klass
	klass = &Class{
		"QStyleOptionTabBarBase",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabBarBase"] = klass
	klass = &Class{
		"QStyleOptionTabBarBaseV2",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabBarBaseV2"] = klass
	klass = &Class{
		"QStyleOptionTabV2",
		false,
		make([]*Class, 0),
		104,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabV2"] = klass
	klass = &Class{
		"QStyleOptionTabV3",
		false,
		make([]*Class, 0),
		128,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabV3"] = klass
	klass = &Class{
		"QStyleOptionTabWidgetFrame",
		false,
		make([]*Class, 0),
		96,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabWidgetFrame"] = klass
	klass = &Class{
		"QStyleOptionTabWidgetFrameV2",
		false,
		make([]*Class, 0),
		128,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTabWidgetFrameV2"] = klass
	klass = &Class{
		"QStyleOptionTitleBar",
		false,
		make([]*Class, 0),
		88,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionTitleBar"] = klass
	klass = &Class{
		"QStyleOptionToolBar",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionToolBar"] = klass
	klass = &Class{
		"QStyleOptionToolBox",
		false,
		make([]*Class, 0),
		72,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionToolBox"] = klass
	klass = &Class{
		"QStyleOptionToolBoxV2",
		false,
		make([]*Class, 0),
		80,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionToolBoxV2"] = klass
	klass = &Class{
		"QStyleOptionToolButton",
		false,
		make([]*Class, 0),
		128,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionToolButton"] = klass
	klass = &Class{
		"QStyleOptionViewItem",
		false,
		make([]*Class, 0),
		104,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionViewItem"] = klass
	klass = &Class{
		"QStyleOptionViewItemV2",
		false,
		make([]*Class, 0),
		104,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionViewItemV2"] = klass
	klass = &Class{
		"QStyleOptionViewItemV3",
		false,
		make([]*Class, 0),
		120,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionViewItemV3"] = klass
	klass = &Class{
		"QStyleOptionViewItemV4",
		false,
		make([]*Class, 0),
		184,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QStyleOptionViewItemV4"] = klass
	klass = &Class{
		"QStylePainter",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QStylePainter"] = klass
	klass = &Class{
		"QStylePlugin",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStylePlugin"] = klass
	klass = &Class{
		"QStyledItemDelegate",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QStyledItemDelegate"] = klass
	klass = &Class{
		"QSwipeGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSwipeGesture"] = klass
	klass = &Class{
		"QSyntaxHighlighter",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSyntaxHighlighter"] = klass
	klass = &Class{
		"QSystemTrayIcon",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QSystemTrayIcon"] = klass
	klass = &Class{
		"QTabBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTabBar"] = klass
	klass = &Class{
		"QTabWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTabWidget"] = klass
	klass = &Class{
		"QTableView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTableView"] = klass
	klass = &Class{
		"QTableWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTableWidget"] = klass
	klass = &Class{
		"QTableWidgetItem",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTableWidgetItem"] = klass
	klass = &Class{
		"QTableWidgetSelectionRange",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTableWidgetSelectionRange"] = klass
	klass = &Class{
		"QTabletEvent",
		false,
		make([]*Class, 0),
		120,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTabletEvent"] = klass
	klass = &Class{
		"QTapAndHoldGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTapAndHoldGesture"] = klass
	klass = &Class{
		"QTapGesture",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTapGesture"] = klass
	klass = &Class{
		"QTextBlock",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextBlock"] = klass
	klass = &Class{
		"QTextBlock::iterator",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextBlock::iterator"] = klass
	klass = &Class{
		"QTextBlockFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextBlockFormat"] = klass
	klass = &Class{
		"QTextBlockGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextBlockGroup"] = klass
	klass = &Class{
		"QTextBlockUserData",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTextBlockUserData"] = klass
	klass = &Class{
		"QTextBrowser",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextBrowser"] = klass
	klass = &Class{
		"QTextCharFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextCharFormat"] = klass
	klass = &Class{
		"QTextCodec",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextCodec"] = klass
	klass = &Class{
		"QTextCursor",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextCursor"] = klass
	klass = &Class{
		"QTextDocument",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextDocument"] = klass
	klass = &Class{
		"QTextDocumentFragment",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextDocumentFragment"] = klass
	klass = &Class{
		"QTextDocumentPrivate",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextDocumentPrivate"] = klass
	klass = &Class{
		"QTextDocumentWriter",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QTextDocumentWriter"] = klass
	klass = &Class{
		"QTextEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextEdit"] = klass
	klass = &Class{
		"QTextEdit::ExtraSelection",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextEdit::ExtraSelection"] = klass
	klass = &Class{
		"QTextEngine",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextEngine"] = klass
	klass = &Class{
		"QTextFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextFormat"] = klass
	klass = &Class{
		"QTextFragment",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextFragment"] = klass
	klass = &Class{
		"QTextFrame",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextFrame"] = klass
	klass = &Class{
		"QTextFrame::iterator",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextFrame::iterator"] = klass
	klass = &Class{
		"QTextFrameFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextFrameFormat"] = klass
	klass = &Class{
		"QTextFrameLayoutData",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTextFrameLayoutData"] = klass
	klass = &Class{
		"QTextImageFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextImageFormat"] = klass
	klass = &Class{
		"QTextInlineObject",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextInlineObject"] = klass
	klass = &Class{
		"QTextItem",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextItem"] = klass
	klass = &Class{
		"QTextLayout",
		false,
		make([]*Class, 0),
		8,
		true,
		false,
		false,
		false,
		false,
	}
	classes["QTextLayout"] = klass
	klass = &Class{
		"QTextLayout::FormatRange",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextLayout::FormatRange"] = klass
	klass = &Class{
		"QTextLength",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextLength"] = klass
	klass = &Class{
		"QTextLine",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextLine"] = klass
	klass = &Class{
		"QTextList",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextList"] = klass
	klass = &Class{
		"QTextListFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextListFormat"] = klass
	klass = &Class{
		"QTextObject",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextObject"] = klass
	klass = &Class{
		"QTextObjectInterface",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTextObjectInterface"] = klass
	klass = &Class{
		"QTextOption",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextOption"] = klass
	klass = &Class{
		"QTextOption::Tab",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextOption::Tab"] = klass
	klass = &Class{
		"QTextStream",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextStream"] = klass
	klass = &Class{
		"QTextStreamManipulator",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTextStreamManipulator"] = klass
	klass = &Class{
		"QTextTable",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTextTable"] = klass
	klass = &Class{
		"QTextTableCell",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextTableCell"] = klass
	klass = &Class{
		"QTextTableCellFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextTableCellFormat"] = klass
	klass = &Class{
		"QTextTableFormat",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTextTableFormat"] = klass
	klass = &Class{
		"QTileRules",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTileRules"] = klass
	klass = &Class{
		"QTime",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTime"] = klass
	klass = &Class{
		"QTimeEdit",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTimeEdit"] = klass
	klass = &Class{
		"QTimeLine",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTimeLine"] = klass
	klass = &Class{
		"QTimerEvent",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QTimerEvent"] = klass
	klass = &Class{
		"QToolBar",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QToolBar"] = klass
	klass = &Class{
		"QToolBarChangeEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QToolBarChangeEvent"] = klass
	klass = &Class{
		"QToolBox",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QToolBox"] = klass
	klass = &Class{
		"QToolButton",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QToolButton"] = klass
	klass = &Class{
		"QToolTip",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QToolTip"] = klass
	klass = &Class{
		"QTouchEvent",
		false,
		make([]*Class, 0),
		48,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTouchEvent"] = klass
	klass = &Class{
		"QTouchEvent::TouchPoint",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTouchEvent::TouchPoint"] = klass
	klass = &Class{
		"QTransform",
		false,
		make([]*Class, 0),
		88,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTransform"] = klass
	klass = &Class{
		"QTreeView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTreeView"] = klass
	klass = &Class{
		"QTreeWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QTreeWidget"] = klass
	klass = &Class{
		"QTreeWidgetItem",
		false,
		make([]*Class, 0),
		64,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QTreeWidgetItem"] = klass
	klass = &Class{
		"QTreeWidgetItemIterator",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QTreeWidgetItemIterator"] = klass
	klass = &Class{
		"QUndoCommand",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QUndoCommand"] = klass
	klass = &Class{
		"QUndoGroup",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QUndoGroup"] = klass
	klass = &Class{
		"QUndoStack",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QUndoStack"] = klass
	klass = &Class{
		"QUndoView",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QUndoView"] = klass
	klass = &Class{
		"QUnixPrintWidget",
		false,
		make([]*Class, 0),
		48,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QUnixPrintWidget"] = klass
	klass = &Class{
		"QUrl",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QUrl"] = klass
	klass = &Class{
		"QUuid",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QUuid"] = klass
	klass = &Class{
		"QVBoxLayout",
		false,
		make([]*Class, 0),
		32,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QVBoxLayout"] = klass
	klass = &Class{
		"QValidator",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QValidator"] = klass
	klass = &Class{
		"QVariant",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QVariant"] = klass
	klass = &Class{
		"QVariantComparisonHelper",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QVariantComparisonHelper"] = klass
	klass = &Class{
		"QVector2D",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QVector2D"] = klass
	klass = &Class{
		"QVector3D",
		false,
		make([]*Class, 0),
		12,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QVector3D"] = klass
	klass = &Class{
		"QVector4D",
		false,
		make([]*Class, 0),
		16,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QVector4D"] = klass
	klass = &Class{
		"QWhatsThis",
		false,
		make([]*Class, 0),
		1,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QWhatsThis"] = klass
	klass = &Class{
		"QWhatsThisClickedEvent",
		false,
		make([]*Class, 0),
		32,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QWhatsThisClickedEvent"] = klass
	klass = &Class{
		"QWheelEvent",
		false,
		make([]*Class, 0),
		56,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QWheelEvent"] = klass
	klass = &Class{
		"QWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWidget"] = klass
	klass = &Class{
		"QWidgetAction",
		false,
		make([]*Class, 0),
		16,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWidgetAction"] = klass
	klass = &Class{
		"QWidgetItem",
		false,
		make([]*Class, 0),
		24,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWidgetItem"] = klass
	klass = &Class{
		"QWidgetItemV2",
		false,
		make([]*Class, 0),
		88,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWidgetItemV2"] = klass
	klass = &Class{
		"QWindowStateChangeEvent",
		false,
		make([]*Class, 0),
		24,
		true,
		true,
		true,
		false,
		false,
	}
	classes["QWindowStateChangeEvent"] = klass
	klass = &Class{
		"QWindowSurface",
		true,
		make([]*Class, 0),
		0,
		false,
		false,
		false,
		false,
		false,
	}
	classes["QWindowSurface"] = klass
	klass = &Class{
		"QWizard",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWizard"] = klass
	klass = &Class{
		"QWizardPage",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWizardPage"] = klass
	klass = &Class{
		"QWorkspace",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QWorkspace"] = klass
	klass = &Class{
		"QX11EmbedContainer",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QX11EmbedContainer"] = klass
	klass = &Class{
		"QX11EmbedWidget",
		false,
		make([]*Class, 0),
		40,
		true,
		false,
		true,
		false,
		false,
	}
	classes["QX11EmbedWidget"] = klass
	klass = &Class{
		"QX11Info",
		false,
		make([]*Class, 0),
		8,
		true,
		true,
		false,
		false,
		false,
	}
	classes["QX11Info"] = klass
	classes["QAbstractButton"].parents = append(classes["QAbstractButton"].parents, classes["QWidget"])
	classes["QAbstractGraphicsShapeItem"].parents = append(classes["QAbstractGraphicsShapeItem"].parents, classes["QGraphicsItem"])
	classes["QAbstractItemDelegate"].parents = append(classes["QAbstractItemDelegate"].parents, classes["QObject"])
	classes["QAbstractItemView"].parents = append(classes["QAbstractItemView"].parents, classes["QAbstractScrollArea"])
	classes["QAbstractPageSetupDialog"].parents = append(classes["QAbstractPageSetupDialog"].parents, classes["QDialog"])
	classes["QAbstractPrintDialog"].parents = append(classes["QAbstractPrintDialog"].parents, classes["QDialog"])
	classes["QAbstractProxyModel"].parents = append(classes["QAbstractProxyModel"].parents, classes["QAbstractItemModel"])
	classes["QAbstractScrollArea"].parents = append(classes["QAbstractScrollArea"].parents, classes["QFrame"])
	classes["QAbstractSlider"].parents = append(classes["QAbstractSlider"].parents, classes["QWidget"])
	classes["QAbstractSpinBox"].parents = append(classes["QAbstractSpinBox"].parents, classes["QWidget"])
	classes["QAbstractTextDocumentLayout"].parents = append(classes["QAbstractTextDocumentLayout"].parents, classes["QObject"])
	classes["QAccessibleActionInterface"].parents = append(classes["QAccessibleActionInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleApplication"].parents = append(classes["QAccessibleApplication"].parents, classes["QAccessibleObject"])
	classes["QAccessibleBridgeFactoryInterface"].parents = append(classes["QAccessibleBridgeFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QAccessibleBridgePlugin"].parents = append(classes["QAccessibleBridgePlugin"].parents, classes["QObject"])
	classes["QAccessibleBridgePlugin"].parents = append(classes["QAccessibleBridgePlugin"].parents, classes["QAccessibleBridgeFactoryInterface"])
	classes["QAccessibleEditableTextInterface"].parents = append(classes["QAccessibleEditableTextInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleEvent"].parents = append(classes["QAccessibleEvent"].parents, classes["QEvent"])
	classes["QAccessibleFactoryInterface"].parents = append(classes["QAccessibleFactoryInterface"].parents, classes["QAccessible"])
	classes["QAccessibleFactoryInterface"].parents = append(classes["QAccessibleFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QAccessibleImageInterface"].parents = append(classes["QAccessibleImageInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleInterface"].parents = append(classes["QAccessibleInterface"].parents, classes["QAccessible"])
	classes["QAccessibleInterfaceEx"].parents = append(classes["QAccessibleInterfaceEx"].parents, classes["QAccessibleInterface"])
	classes["QAccessibleObject"].parents = append(classes["QAccessibleObject"].parents, classes["QAccessibleInterface"])
	classes["QAccessibleObjectEx"].parents = append(classes["QAccessibleObjectEx"].parents, classes["QAccessibleInterfaceEx"])
	classes["QAccessiblePlugin"].parents = append(classes["QAccessiblePlugin"].parents, classes["QObject"])
	classes["QAccessiblePlugin"].parents = append(classes["QAccessiblePlugin"].parents, classes["QAccessibleFactoryInterface"])
	classes["QAccessibleSimpleEditableTextInterface"].parents = append(classes["QAccessibleSimpleEditableTextInterface"].parents, classes["QAccessibleEditableTextInterface"])
	classes["QAccessibleTable2CellInterface"].parents = append(classes["QAccessibleTable2CellInterface"].parents, classes["QAccessibleInterface"])
	classes["QAccessibleTable2Interface"].parents = append(classes["QAccessibleTable2Interface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleTableInterface"].parents = append(classes["QAccessibleTableInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleTextInterface"].parents = append(classes["QAccessibleTextInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleValueInterface"].parents = append(classes["QAccessibleValueInterface"].parents, classes["QAccessible2Interface"])
	classes["QAccessibleWidget"].parents = append(classes["QAccessibleWidget"].parents, classes["QAccessibleObject"])
	classes["QAccessibleWidgetEx"].parents = append(classes["QAccessibleWidgetEx"].parents, classes["QAccessibleObjectEx"])
	classes["QAction"].parents = append(classes["QAction"].parents, classes["QObject"])
	classes["QActionEvent"].parents = append(classes["QActionEvent"].parents, classes["QEvent"])
	classes["QActionGroup"].parents = append(classes["QActionGroup"].parents, classes["QObject"])
	classes["QApplication"].parents = append(classes["QApplication"].parents, classes["QCoreApplication"])
	classes["QBitmap"].parents = append(classes["QBitmap"].parents, classes["QPixmap"])
	classes["QBoxLayout"].parents = append(classes["QBoxLayout"].parents, classes["QLayout"])
	classes["QButtonGroup"].parents = append(classes["QButtonGroup"].parents, classes["QObject"])
	classes["QCalendarWidget"].parents = append(classes["QCalendarWidget"].parents, classes["QWidget"])
	classes["QCheckBox"].parents = append(classes["QCheckBox"].parents, classes["QAbstractButton"])
	classes["QClipboard"].parents = append(classes["QClipboard"].parents, classes["QObject"])
	classes["QClipboardEvent"].parents = append(classes["QClipboardEvent"].parents, classes["QEvent"])
	classes["QCloseEvent"].parents = append(classes["QCloseEvent"].parents, classes["QEvent"])
	classes["QColorDialog"].parents = append(classes["QColorDialog"].parents, classes["QDialog"])
	classes["QColumnView"].parents = append(classes["QColumnView"].parents, classes["QAbstractItemView"])
	classes["QComboBox"].parents = append(classes["QComboBox"].parents, classes["QWidget"])
	classes["QCommandLinkButton"].parents = append(classes["QCommandLinkButton"].parents, classes["QPushButton"])
	classes["QCommonStyle"].parents = append(classes["QCommonStyle"].parents, classes["QStyle"])
	classes["QCompleter"].parents = append(classes["QCompleter"].parents, classes["QObject"])
	classes["QConicalGradient"].parents = append(classes["QConicalGradient"].parents, classes["QGradient"])
	classes["QContextMenuEvent"].parents = append(classes["QContextMenuEvent"].parents, classes["QInputEvent"])
	classes["QDataWidgetMapper"].parents = append(classes["QDataWidgetMapper"].parents, classes["QObject"])
	classes["QDateEdit"].parents = append(classes["QDateEdit"].parents, classes["QDateTimeEdit"])
	classes["QDateTimeEdit"].parents = append(classes["QDateTimeEdit"].parents, classes["QAbstractSpinBox"])
	classes["QDesktopWidget"].parents = append(classes["QDesktopWidget"].parents, classes["QWidget"])
	classes["QDial"].parents = append(classes["QDial"].parents, classes["QAbstractSlider"])
	classes["QDialog"].parents = append(classes["QDialog"].parents, classes["QWidget"])
	classes["QDialogButtonBox"].parents = append(classes["QDialogButtonBox"].parents, classes["QWidget"])
	classes["QDirModel"].parents = append(classes["QDirModel"].parents, classes["QAbstractItemModel"])
	classes["QDockWidget"].parents = append(classes["QDockWidget"].parents, classes["QWidget"])
	classes["QDoubleSpinBox"].parents = append(classes["QDoubleSpinBox"].parents, classes["QAbstractSpinBox"])
	classes["QDoubleValidator"].parents = append(classes["QDoubleValidator"].parents, classes["QValidator"])
	classes["QDrag"].parents = append(classes["QDrag"].parents, classes["QObject"])
	classes["QDragEnterEvent"].parents = append(classes["QDragEnterEvent"].parents, classes["QDragMoveEvent"])
	classes["QDragLeaveEvent"].parents = append(classes["QDragLeaveEvent"].parents, classes["QEvent"])
	classes["QDragMoveEvent"].parents = append(classes["QDragMoveEvent"].parents, classes["QDropEvent"])
	classes["QDragResponseEvent"].parents = append(classes["QDragResponseEvent"].parents, classes["QEvent"])
	classes["QDropEvent"].parents = append(classes["QDropEvent"].parents, classes["QEvent"])
	classes["QDropEvent"].parents = append(classes["QDropEvent"].parents, classes["QMimeSource"])
	classes["QErrorMessage"].parents = append(classes["QErrorMessage"].parents, classes["QDialog"])
	classes["QFileDialog"].parents = append(classes["QFileDialog"].parents, classes["QDialog"])
	classes["QFileOpenEvent"].parents = append(classes["QFileOpenEvent"].parents, classes["QEvent"])
	classes["QFileSystemModel"].parents = append(classes["QFileSystemModel"].parents, classes["QAbstractItemModel"])
	classes["QFocusEvent"].parents = append(classes["QFocusEvent"].parents, classes["QEvent"])
	classes["QFocusFrame"].parents = append(classes["QFocusFrame"].parents, classes["QWidget"])
	classes["QFontComboBox"].parents = append(classes["QFontComboBox"].parents, classes["QComboBox"])
	classes["QFontDialog"].parents = append(classes["QFontDialog"].parents, classes["QDialog"])
	classes["QFormLayout"].parents = append(classes["QFormLayout"].parents, classes["QLayout"])
	classes["QFrame"].parents = append(classes["QFrame"].parents, classes["QWidget"])
	classes["QGesture"].parents = append(classes["QGesture"].parents, classes["QObject"])
	classes["QGestureEvent"].parents = append(classes["QGestureEvent"].parents, classes["QEvent"])
	classes["QGraphicsAnchor"].parents = append(classes["QGraphicsAnchor"].parents, classes["QObject"])
	classes["QGraphicsAnchorLayout"].parents = append(classes["QGraphicsAnchorLayout"].parents, classes["QGraphicsLayout"])
	classes["QGraphicsBlurEffect"].parents = append(classes["QGraphicsBlurEffect"].parents, classes["QGraphicsEffect"])
	classes["QGraphicsColorizeEffect"].parents = append(classes["QGraphicsColorizeEffect"].parents, classes["QGraphicsEffect"])
	classes["QGraphicsDropShadowEffect"].parents = append(classes["QGraphicsDropShadowEffect"].parents, classes["QGraphicsEffect"])
	classes["QGraphicsEffect"].parents = append(classes["QGraphicsEffect"].parents, classes["QObject"])
	classes["QGraphicsEllipseItem"].parents = append(classes["QGraphicsEllipseItem"].parents, classes["QAbstractGraphicsShapeItem"])
	classes["QGraphicsGridLayout"].parents = append(classes["QGraphicsGridLayout"].parents, classes["QGraphicsLayout"])
	classes["QGraphicsItemAnimation"].parents = append(classes["QGraphicsItemAnimation"].parents, classes["QObject"])
	classes["QGraphicsItemGroup"].parents = append(classes["QGraphicsItemGroup"].parents, classes["QGraphicsItem"])
	classes["QGraphicsLayout"].parents = append(classes["QGraphicsLayout"].parents, classes["QGraphicsLayoutItem"])
	classes["QGraphicsLineItem"].parents = append(classes["QGraphicsLineItem"].parents, classes["QGraphicsItem"])
	classes["QGraphicsLinearLayout"].parents = append(classes["QGraphicsLinearLayout"].parents, classes["QGraphicsLayout"])
	classes["QGraphicsObject"].parents = append(classes["QGraphicsObject"].parents, classes["QObject"])
	classes["QGraphicsObject"].parents = append(classes["QGraphicsObject"].parents, classes["QGraphicsItem"])
	classes["QGraphicsOpacityEffect"].parents = append(classes["QGraphicsOpacityEffect"].parents, classes["QGraphicsEffect"])
	classes["QGraphicsPathItem"].parents = append(classes["QGraphicsPathItem"].parents, classes["QAbstractGraphicsShapeItem"])
	classes["QGraphicsPixmapItem"].parents = append(classes["QGraphicsPixmapItem"].parents, classes["QGraphicsItem"])
	classes["QGraphicsPolygonItem"].parents = append(classes["QGraphicsPolygonItem"].parents, classes["QAbstractGraphicsShapeItem"])
	classes["QGraphicsProxyWidget"].parents = append(classes["QGraphicsProxyWidget"].parents, classes["QGraphicsWidget"])
	classes["QGraphicsRectItem"].parents = append(classes["QGraphicsRectItem"].parents, classes["QAbstractGraphicsShapeItem"])
	classes["QGraphicsRotation"].parents = append(classes["QGraphicsRotation"].parents, classes["QGraphicsTransform"])
	classes["QGraphicsScale"].parents = append(classes["QGraphicsScale"].parents, classes["QGraphicsTransform"])
	classes["QGraphicsScene"].parents = append(classes["QGraphicsScene"].parents, classes["QObject"])
	classes["QGraphicsSceneContextMenuEvent"].parents = append(classes["QGraphicsSceneContextMenuEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneDragDropEvent"].parents = append(classes["QGraphicsSceneDragDropEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneEvent"].parents = append(classes["QGraphicsSceneEvent"].parents, classes["QEvent"])
	classes["QGraphicsSceneHelpEvent"].parents = append(classes["QGraphicsSceneHelpEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneHoverEvent"].parents = append(classes["QGraphicsSceneHoverEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneMouseEvent"].parents = append(classes["QGraphicsSceneMouseEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneMoveEvent"].parents = append(classes["QGraphicsSceneMoveEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneResizeEvent"].parents = append(classes["QGraphicsSceneResizeEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSceneWheelEvent"].parents = append(classes["QGraphicsSceneWheelEvent"].parents, classes["QGraphicsSceneEvent"])
	classes["QGraphicsSimpleTextItem"].parents = append(classes["QGraphicsSimpleTextItem"].parents, classes["QAbstractGraphicsShapeItem"])
	classes["QGraphicsTextItem"].parents = append(classes["QGraphicsTextItem"].parents, classes["QGraphicsObject"])
	classes["QGraphicsTransform"].parents = append(classes["QGraphicsTransform"].parents, classes["QObject"])
	classes["QGraphicsView"].parents = append(classes["QGraphicsView"].parents, classes["QAbstractScrollArea"])
	classes["QGraphicsWidget"].parents = append(classes["QGraphicsWidget"].parents, classes["QGraphicsObject"])
	classes["QGraphicsWidget"].parents = append(classes["QGraphicsWidget"].parents, classes["QGraphicsLayoutItem"])
	classes["QGridLayout"].parents = append(classes["QGridLayout"].parents, classes["QLayout"])
	classes["QGroupBox"].parents = append(classes["QGroupBox"].parents, classes["QWidget"])
	classes["QHBoxLayout"].parents = append(classes["QHBoxLayout"].parents, classes["QBoxLayout"])
	classes["QHeaderView"].parents = append(classes["QHeaderView"].parents, classes["QAbstractItemView"])
	classes["QHelpEvent"].parents = append(classes["QHelpEvent"].parents, classes["QEvent"])
	classes["QHideEvent"].parents = append(classes["QHideEvent"].parents, classes["QEvent"])
	classes["QHoverEvent"].parents = append(classes["QHoverEvent"].parents, classes["QEvent"])
	classes["QIconDragEvent"].parents = append(classes["QIconDragEvent"].parents, classes["QEvent"])
	classes["QIconEngineFactoryInterface"].parents = append(classes["QIconEngineFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QIconEngineFactoryInterfaceV2"].parents = append(classes["QIconEngineFactoryInterfaceV2"].parents, classes["QFactoryInterface"])
	classes["QIconEnginePlugin"].parents = append(classes["QIconEnginePlugin"].parents, classes["QObject"])
	classes["QIconEnginePlugin"].parents = append(classes["QIconEnginePlugin"].parents, classes["QIconEngineFactoryInterface"])
	classes["QIconEnginePluginV2"].parents = append(classes["QIconEnginePluginV2"].parents, classes["QObject"])
	classes["QIconEnginePluginV2"].parents = append(classes["QIconEnginePluginV2"].parents, classes["QIconEngineFactoryInterfaceV2"])
	classes["QIconEngineV2"].parents = append(classes["QIconEngineV2"].parents, classes["QIconEngine"])
	classes["QImage"].parents = append(classes["QImage"].parents, classes["QPaintDevice"])
	classes["QImageIOHandlerFactoryInterface"].parents = append(classes["QImageIOHandlerFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QImageIOPlugin"].parents = append(classes["QImageIOPlugin"].parents, classes["QObject"])
	classes["QImageIOPlugin"].parents = append(classes["QImageIOPlugin"].parents, classes["QImageIOHandlerFactoryInterface"])
	classes["QInputContext"].parents = append(classes["QInputContext"].parents, classes["QObject"])
	classes["QInputContextFactoryInterface"].parents = append(classes["QInputContextFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QInputContextPlugin"].parents = append(classes["QInputContextPlugin"].parents, classes["QObject"])
	classes["QInputContextPlugin"].parents = append(classes["QInputContextPlugin"].parents, classes["QInputContextFactoryInterface"])
	classes["QInputDialog"].parents = append(classes["QInputDialog"].parents, classes["QDialog"])
	classes["QInputEvent"].parents = append(classes["QInputEvent"].parents, classes["QEvent"])
	classes["QInputMethodEvent"].parents = append(classes["QInputMethodEvent"].parents, classes["QEvent"])
	classes["QIntValidator"].parents = append(classes["QIntValidator"].parents, classes["QValidator"])
	classes["QItemDelegate"].parents = append(classes["QItemDelegate"].parents, classes["QAbstractItemDelegate"])
	classes["QItemSelectionModel"].parents = append(classes["QItemSelectionModel"].parents, classes["QObject"])
	classes["QKeyEvent"].parents = append(classes["QKeyEvent"].parents, classes["QInputEvent"])
	classes["QKeyEventTransition"].parents = append(classes["QKeyEventTransition"].parents, classes["QEventTransition"])
	classes["QLCDNumber"].parents = append(classes["QLCDNumber"].parents, classes["QFrame"])
	classes["QLabel"].parents = append(classes["QLabel"].parents, classes["QFrame"])
	classes["QLayout"].parents = append(classes["QLayout"].parents, classes["QObject"])
	classes["QLayout"].parents = append(classes["QLayout"].parents, classes["QLayoutItem"])
	classes["QLineEdit"].parents = append(classes["QLineEdit"].parents, classes["QWidget"])
	classes["QLinearGradient"].parents = append(classes["QLinearGradient"].parents, classes["QGradient"])
	classes["QListView"].parents = append(classes["QListView"].parents, classes["QAbstractItemView"])
	classes["QListWidget"].parents = append(classes["QListWidget"].parents, classes["QListView"])
	classes["QMainWindow"].parents = append(classes["QMainWindow"].parents, classes["QWidget"])
	classes["QMdiArea"].parents = append(classes["QMdiArea"].parents, classes["QAbstractScrollArea"])
	classes["QMdiSubWindow"].parents = append(classes["QMdiSubWindow"].parents, classes["QWidget"])
	classes["QMenu"].parents = append(classes["QMenu"].parents, classes["QWidget"])
	classes["QMenuBar"].parents = append(classes["QMenuBar"].parents, classes["QWidget"])
	classes["QMessageBox"].parents = append(classes["QMessageBox"].parents, classes["QDialog"])
	classes["QMouseEvent"].parents = append(classes["QMouseEvent"].parents, classes["QInputEvent"])
	classes["QMouseEventTransition"].parents = append(classes["QMouseEventTransition"].parents, classes["QEventTransition"])
	classes["QMoveEvent"].parents = append(classes["QMoveEvent"].parents, classes["QEvent"])
	classes["QMovie"].parents = append(classes["QMovie"].parents, classes["QObject"])
	classes["QPageSetupDialog"].parents = append(classes["QPageSetupDialog"].parents, classes["QAbstractPageSetupDialog"])
	classes["QPaintEvent"].parents = append(classes["QPaintEvent"].parents, classes["QEvent"])
	classes["QPanGesture"].parents = append(classes["QPanGesture"].parents, classes["QGesture"])
	classes["QPicture"].parents = append(classes["QPicture"].parents, classes["QPaintDevice"])
	classes["QPictureFormatInterface"].parents = append(classes["QPictureFormatInterface"].parents, classes["QFactoryInterface"])
	classes["QPictureFormatPlugin"].parents = append(classes["QPictureFormatPlugin"].parents, classes["QObject"])
	classes["QPictureFormatPlugin"].parents = append(classes["QPictureFormatPlugin"].parents, classes["QPictureFormatInterface"])
	classes["QPinchGesture"].parents = append(classes["QPinchGesture"].parents, classes["QGesture"])
	classes["QPixmap"].parents = append(classes["QPixmap"].parents, classes["QPaintDevice"])
	classes["QPlainTextDocumentLayout"].parents = append(classes["QPlainTextDocumentLayout"].parents, classes["QAbstractTextDocumentLayout"])
	classes["QPlainTextEdit"].parents = append(classes["QPlainTextEdit"].parents, classes["QAbstractScrollArea"])
	classes["QPrintDialog"].parents = append(classes["QPrintDialog"].parents, classes["QAbstractPrintDialog"])
	classes["QPrintPreviewDialog"].parents = append(classes["QPrintPreviewDialog"].parents, classes["QDialog"])
	classes["QPrintPreviewWidget"].parents = append(classes["QPrintPreviewWidget"].parents, classes["QWidget"])
	classes["QPrinter"].parents = append(classes["QPrinter"].parents, classes["QPaintDevice"])
	classes["QProgressBar"].parents = append(classes["QProgressBar"].parents, classes["QWidget"])
	classes["QProgressDialog"].parents = append(classes["QProgressDialog"].parents, classes["QDialog"])
	classes["QProxyModel"].parents = append(classes["QProxyModel"].parents, classes["QAbstractItemModel"])
	classes["QProxyStyle"].parents = append(classes["QProxyStyle"].parents, classes["QCommonStyle"])
	classes["QPushButton"].parents = append(classes["QPushButton"].parents, classes["QAbstractButton"])
	classes["QRadialGradient"].parents = append(classes["QRadialGradient"].parents, classes["QGradient"])
	classes["QRadioButton"].parents = append(classes["QRadioButton"].parents, classes["QAbstractButton"])
	classes["QRegExpValidator"].parents = append(classes["QRegExpValidator"].parents, classes["QValidator"])
	classes["QResizeEvent"].parents = append(classes["QResizeEvent"].parents, classes["QEvent"])
	classes["QRubberBand"].parents = append(classes["QRubberBand"].parents, classes["QWidget"])
	classes["QScrollArea"].parents = append(classes["QScrollArea"].parents, classes["QAbstractScrollArea"])
	classes["QScrollBar"].parents = append(classes["QScrollBar"].parents, classes["QAbstractSlider"])
	classes["QSessionManager"].parents = append(classes["QSessionManager"].parents, classes["QObject"])
	classes["QShortcut"].parents = append(classes["QShortcut"].parents, classes["QObject"])
	classes["QShortcutEvent"].parents = append(classes["QShortcutEvent"].parents, classes["QEvent"])
	classes["QShowEvent"].parents = append(classes["QShowEvent"].parents, classes["QEvent"])
	classes["QSizeGrip"].parents = append(classes["QSizeGrip"].parents, classes["QWidget"])
	classes["QSlider"].parents = append(classes["QSlider"].parents, classes["QAbstractSlider"])
	classes["QSortFilterProxyModel"].parents = append(classes["QSortFilterProxyModel"].parents, classes["QAbstractProxyModel"])
	classes["QSound"].parents = append(classes["QSound"].parents, classes["QObject"])
	classes["QSpacerItem"].parents = append(classes["QSpacerItem"].parents, classes["QLayoutItem"])
	classes["QSpinBox"].parents = append(classes["QSpinBox"].parents, classes["QAbstractSpinBox"])
	classes["QSplashScreen"].parents = append(classes["QSplashScreen"].parents, classes["QWidget"])
	classes["QSplitter"].parents = append(classes["QSplitter"].parents, classes["QFrame"])
	classes["QSplitterHandle"].parents = append(classes["QSplitterHandle"].parents, classes["QWidget"])
	classes["QStackedLayout"].parents = append(classes["QStackedLayout"].parents, classes["QLayout"])
	classes["QStackedWidget"].parents = append(classes["QStackedWidget"].parents, classes["QFrame"])
	classes["QStandardItemModel"].parents = append(classes["QStandardItemModel"].parents, classes["QAbstractItemModel"])
	classes["QStatusBar"].parents = append(classes["QStatusBar"].parents, classes["QWidget"])
	classes["QStatusTipEvent"].parents = append(classes["QStatusTipEvent"].parents, classes["QEvent"])
	classes["QStringListModel"].parents = append(classes["QStringListModel"].parents, classes["QAbstractListModel"])
	classes["QStyle"].parents = append(classes["QStyle"].parents, classes["QObject"])
	classes["QStyleFactoryInterface"].parents = append(classes["QStyleFactoryInterface"].parents, classes["QFactoryInterface"])
	classes["QStyleHintReturnMask"].parents = append(classes["QStyleHintReturnMask"].parents, classes["QStyleHintReturn"])
	classes["QStyleHintReturnVariant"].parents = append(classes["QStyleHintReturnVariant"].parents, classes["QStyleHintReturn"])
	classes["QStyleOptionButton"].parents = append(classes["QStyleOptionButton"].parents, classes["QStyleOption"])
	classes["QStyleOptionComboBox"].parents = append(classes["QStyleOptionComboBox"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionComplex"].parents = append(classes["QStyleOptionComplex"].parents, classes["QStyleOption"])
	classes["QStyleOptionDockWidget"].parents = append(classes["QStyleOptionDockWidget"].parents, classes["QStyleOption"])
	classes["QStyleOptionDockWidgetV2"].parents = append(classes["QStyleOptionDockWidgetV2"].parents, classes["QStyleOptionDockWidget"])
	classes["QStyleOptionFocusRect"].parents = append(classes["QStyleOptionFocusRect"].parents, classes["QStyleOption"])
	classes["QStyleOptionFrame"].parents = append(classes["QStyleOptionFrame"].parents, classes["QStyleOption"])
	classes["QStyleOptionFrameV2"].parents = append(classes["QStyleOptionFrameV2"].parents, classes["QStyleOptionFrame"])
	classes["QStyleOptionFrameV3"].parents = append(classes["QStyleOptionFrameV3"].parents, classes["QStyleOptionFrameV2"])
	classes["QStyleOptionGraphicsItem"].parents = append(classes["QStyleOptionGraphicsItem"].parents, classes["QStyleOption"])
	classes["QStyleOptionGroupBox"].parents = append(classes["QStyleOptionGroupBox"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionHeader"].parents = append(classes["QStyleOptionHeader"].parents, classes["QStyleOption"])
	classes["QStyleOptionMenuItem"].parents = append(classes["QStyleOptionMenuItem"].parents, classes["QStyleOption"])
	classes["QStyleOptionProgressBar"].parents = append(classes["QStyleOptionProgressBar"].parents, classes["QStyleOption"])
	classes["QStyleOptionProgressBarV2"].parents = append(classes["QStyleOptionProgressBarV2"].parents, classes["QStyleOptionProgressBar"])
	classes["QStyleOptionRubberBand"].parents = append(classes["QStyleOptionRubberBand"].parents, classes["QStyleOption"])
	classes["QStyleOptionSizeGrip"].parents = append(classes["QStyleOptionSizeGrip"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionSlider"].parents = append(classes["QStyleOptionSlider"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionSpinBox"].parents = append(classes["QStyleOptionSpinBox"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionTab"].parents = append(classes["QStyleOptionTab"].parents, classes["QStyleOption"])
	classes["QStyleOptionTabBarBase"].parents = append(classes["QStyleOptionTabBarBase"].parents, classes["QStyleOption"])
	classes["QStyleOptionTabBarBaseV2"].parents = append(classes["QStyleOptionTabBarBaseV2"].parents, classes["QStyleOptionTabBarBase"])
	classes["QStyleOptionTabV2"].parents = append(classes["QStyleOptionTabV2"].parents, classes["QStyleOptionTab"])
	classes["QStyleOptionTabV3"].parents = append(classes["QStyleOptionTabV3"].parents, classes["QStyleOptionTabV2"])
	classes["QStyleOptionTabWidgetFrame"].parents = append(classes["QStyleOptionTabWidgetFrame"].parents, classes["QStyleOption"])
	classes["QStyleOptionTabWidgetFrameV2"].parents = append(classes["QStyleOptionTabWidgetFrameV2"].parents, classes["QStyleOptionTabWidgetFrame"])
	classes["QStyleOptionTitleBar"].parents = append(classes["QStyleOptionTitleBar"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionToolBar"].parents = append(classes["QStyleOptionToolBar"].parents, classes["QStyleOption"])
	classes["QStyleOptionToolBox"].parents = append(classes["QStyleOptionToolBox"].parents, classes["QStyleOption"])
	classes["QStyleOptionToolBoxV2"].parents = append(classes["QStyleOptionToolBoxV2"].parents, classes["QStyleOptionToolBox"])
	classes["QStyleOptionToolButton"].parents = append(classes["QStyleOptionToolButton"].parents, classes["QStyleOptionComplex"])
	classes["QStyleOptionViewItem"].parents = append(classes["QStyleOptionViewItem"].parents, classes["QStyleOption"])
	classes["QStyleOptionViewItemV2"].parents = append(classes["QStyleOptionViewItemV2"].parents, classes["QStyleOptionViewItem"])
	classes["QStyleOptionViewItemV3"].parents = append(classes["QStyleOptionViewItemV3"].parents, classes["QStyleOptionViewItemV2"])
	classes["QStyleOptionViewItemV4"].parents = append(classes["QStyleOptionViewItemV4"].parents, classes["QStyleOptionViewItemV3"])
	classes["QStylePainter"].parents = append(classes["QStylePainter"].parents, classes["QPainter"])
	classes["QStylePlugin"].parents = append(classes["QStylePlugin"].parents, classes["QObject"])
	classes["QStylePlugin"].parents = append(classes["QStylePlugin"].parents, classes["QStyleFactoryInterface"])
	classes["QStyledItemDelegate"].parents = append(classes["QStyledItemDelegate"].parents, classes["QAbstractItemDelegate"])
	classes["QSwipeGesture"].parents = append(classes["QSwipeGesture"].parents, classes["QGesture"])
	classes["QSyntaxHighlighter"].parents = append(classes["QSyntaxHighlighter"].parents, classes["QObject"])
	classes["QSystemTrayIcon"].parents = append(classes["QSystemTrayIcon"].parents, classes["QObject"])
	classes["QTabBar"].parents = append(classes["QTabBar"].parents, classes["QWidget"])
	classes["QTabWidget"].parents = append(classes["QTabWidget"].parents, classes["QWidget"])
	classes["QTableView"].parents = append(classes["QTableView"].parents, classes["QAbstractItemView"])
	classes["QTableWidget"].parents = append(classes["QTableWidget"].parents, classes["QTableView"])
	classes["QTabletEvent"].parents = append(classes["QTabletEvent"].parents, classes["QInputEvent"])
	classes["QTapAndHoldGesture"].parents = append(classes["QTapAndHoldGesture"].parents, classes["QGesture"])
	classes["QTapGesture"].parents = append(classes["QTapGesture"].parents, classes["QGesture"])
	classes["QTextBlockFormat"].parents = append(classes["QTextBlockFormat"].parents, classes["QTextFormat"])
	classes["QTextBlockGroup"].parents = append(classes["QTextBlockGroup"].parents, classes["QTextObject"])
	classes["QTextBrowser"].parents = append(classes["QTextBrowser"].parents, classes["QTextEdit"])
	classes["QTextCharFormat"].parents = append(classes["QTextCharFormat"].parents, classes["QTextFormat"])
	classes["QTextDocument"].parents = append(classes["QTextDocument"].parents, classes["QObject"])
	classes["QTextEdit"].parents = append(classes["QTextEdit"].parents, classes["QAbstractScrollArea"])
	classes["QTextFrame"].parents = append(classes["QTextFrame"].parents, classes["QTextObject"])
	classes["QTextFrameFormat"].parents = append(classes["QTextFrameFormat"].parents, classes["QTextFormat"])
	classes["QTextImageFormat"].parents = append(classes["QTextImageFormat"].parents, classes["QTextCharFormat"])
	classes["QTextList"].parents = append(classes["QTextList"].parents, classes["QTextBlockGroup"])
	classes["QTextListFormat"].parents = append(classes["QTextListFormat"].parents, classes["QTextFormat"])
	classes["QTextObject"].parents = append(classes["QTextObject"].parents, classes["QObject"])
	classes["QTextTable"].parents = append(classes["QTextTable"].parents, classes["QTextFrame"])
	classes["QTextTableCellFormat"].parents = append(classes["QTextTableCellFormat"].parents, classes["QTextCharFormat"])
	classes["QTextTableFormat"].parents = append(classes["QTextTableFormat"].parents, classes["QTextFrameFormat"])
	classes["QTimeEdit"].parents = append(classes["QTimeEdit"].parents, classes["QDateTimeEdit"])
	classes["QToolBar"].parents = append(classes["QToolBar"].parents, classes["QWidget"])
	classes["QToolBarChangeEvent"].parents = append(classes["QToolBarChangeEvent"].parents, classes["QEvent"])
	classes["QToolBox"].parents = append(classes["QToolBox"].parents, classes["QFrame"])
	classes["QToolButton"].parents = append(classes["QToolButton"].parents, classes["QAbstractButton"])
	classes["QTouchEvent"].parents = append(classes["QTouchEvent"].parents, classes["QInputEvent"])
	classes["QTreeView"].parents = append(classes["QTreeView"].parents, classes["QAbstractItemView"])
	classes["QTreeWidget"].parents = append(classes["QTreeWidget"].parents, classes["QTreeView"])
	classes["QUndoGroup"].parents = append(classes["QUndoGroup"].parents, classes["QObject"])
	classes["QUndoStack"].parents = append(classes["QUndoStack"].parents, classes["QObject"])
	classes["QUndoView"].parents = append(classes["QUndoView"].parents, classes["QListView"])
	classes["QUnixPrintWidget"].parents = append(classes["QUnixPrintWidget"].parents, classes["QWidget"])
	classes["QVBoxLayout"].parents = append(classes["QVBoxLayout"].parents, classes["QBoxLayout"])
	classes["QValidator"].parents = append(classes["QValidator"].parents, classes["QObject"])
	classes["QWhatsThisClickedEvent"].parents = append(classes["QWhatsThisClickedEvent"].parents, classes["QEvent"])
	classes["QWheelEvent"].parents = append(classes["QWheelEvent"].parents, classes["QInputEvent"])
	classes["QWidget"].parents = append(classes["QWidget"].parents, classes["QObject"])
	classes["QWidget"].parents = append(classes["QWidget"].parents, classes["QPaintDevice"])
	classes["QWidgetAction"].parents = append(classes["QWidgetAction"].parents, classes["QAction"])
	classes["QWidgetItem"].parents = append(classes["QWidgetItem"].parents, classes["QLayoutItem"])
	classes["QWidgetItemV2"].parents = append(classes["QWidgetItemV2"].parents, classes["QWidgetItem"])
	classes["QWindowStateChangeEvent"].parents = append(classes["QWindowStateChangeEvent"].parents, classes["QEvent"])
	classes["QWizard"].parents = append(classes["QWizard"].parents, classes["QDialog"])
	classes["QWizardPage"].parents = append(classes["QWizardPage"].parents, classes["QWidget"])
	classes["QWorkspace"].parents = append(classes["QWorkspace"].parents, classes["QWidget"])
	classes["QX11EmbedContainer"].parents = append(classes["QX11EmbedContainer"].parents, classes["QWidget"])
	classes["QX11EmbedWidget"].parents = append(classes["QX11EmbedWidget"].parents, classes["QWidget"])
	type_ = &Type{
		"FT_FaceRec_*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1] = type_
	type_ = &Type{
		"QAbstractButton*",
		classes["QAbstractButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[2] = type_
	type_ = &Type{
		"QAbstractFileEngine::FileFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[3] = type_
	type_ = &Type{
		"QAbstractGraphicsShapeItem*",
		classes["QAbstractGraphicsShapeItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[4] = type_
	type_ = &Type{
		"QAbstractItemDelegate*",
		classes["QAbstractItemDelegate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[5] = type_
	type_ = &Type{
		"QAbstractItemDelegate::EndEditHint",
		classes["QAbstractItemDelegate"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[6] = type_
	type_ = &Type{
		"QAbstractItemModel*",
		classes["QAbstractItemModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[7] = type_
	type_ = &Type{
		"QAbstractItemView*",
		classes["QAbstractItemView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[8] = type_
	type_ = &Type{
		"QAbstractItemView::CursorAction",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[9] = type_
	type_ = &Type{
		"QAbstractItemView::DragDropMode",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[10] = type_
	type_ = &Type{
		"QAbstractItemView::DropIndicatorPosition",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[11] = type_
	type_ = &Type{
		"QAbstractItemView::EditTrigger",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[12] = type_
	type_ = &Type{
		"QAbstractItemView::ScrollHint",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[13] = type_
	type_ = &Type{
		"QAbstractItemView::ScrollMode",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[14] = type_
	type_ = &Type{
		"QAbstractItemView::SelectionBehavior",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[15] = type_
	type_ = &Type{
		"QAbstractItemView::SelectionMode",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[16] = type_
	type_ = &Type{
		"QAbstractItemView::State",
		classes["QAbstractItemView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[17] = type_
	type_ = &Type{
		"QAbstractPageSetupDialog*",
		classes["QAbstractPageSetupDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[18] = type_
	type_ = &Type{
		"QAbstractPrintDialog*",
		classes["QAbstractPrintDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[19] = type_
	type_ = &Type{
		"QAbstractPrintDialog::PrintDialogOption",
		classes["QAbstractPrintDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[20] = type_
	type_ = &Type{
		"QAbstractPrintDialog::PrintRange",
		classes["QAbstractPrintDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[21] = type_
	type_ = &Type{
		"QAbstractProxyModel*",
		classes["QAbstractProxyModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[22] = type_
	type_ = &Type{
		"QAbstractScrollArea*",
		classes["QAbstractScrollArea"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[23] = type_
	type_ = &Type{
		"QAbstractSlider*",
		classes["QAbstractSlider"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[24] = type_
	type_ = &Type{
		"QAbstractSlider::SliderAction",
		classes["QAbstractSlider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[25] = type_
	type_ = &Type{
		"QAbstractSlider::SliderChange",
		classes["QAbstractSlider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[26] = type_
	type_ = &Type{
		"QAbstractSpinBox*",
		classes["QAbstractSpinBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[27] = type_
	type_ = &Type{
		"QAbstractSpinBox::ButtonSymbols",
		classes["QAbstractSpinBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[28] = type_
	type_ = &Type{
		"QAbstractSpinBox::CorrectionMode",
		classes["QAbstractSpinBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[29] = type_
	type_ = &Type{
		"QAbstractSpinBox::StepEnabledFlag",
		classes["QAbstractSpinBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[30] = type_
	type_ = &Type{
		"QAbstractTextDocumentLayout*",
		classes["QAbstractTextDocumentLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[31] = type_
	type_ = &Type{
		"QAbstractTextDocumentLayout::PaintContext",
		classes["QAbstractTextDocumentLayout::PaintContext"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[32] = type_
	type_ = &Type{
		"QAbstractTextDocumentLayout::PaintContext*",
		classes["QAbstractTextDocumentLayout::PaintContext"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[33] = type_
	type_ = &Type{
		"QAbstractTextDocumentLayout::Selection*",
		classes["QAbstractTextDocumentLayout::Selection"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[34] = type_
	type_ = &Type{
		"QAbstractUndoItem*",
		classes["QAbstractUndoItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[35] = type_
	type_ = &Type{
		"QAccessible*",
		classes["QAccessible"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[36] = type_
	type_ = &Type{
		"QAccessible2::BoundaryType",
		classes["QAccessible2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[37] = type_
	type_ = &Type{
		"QAccessible2::CoordinateType",
		classes["QAccessible2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[38] = type_
	type_ = &Type{
		"QAccessible2::InterfaceType",
		classes["QAccessible2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[39] = type_
	type_ = &Type{
		"QAccessible2::TableModelChange",
		classes["QAccessible2::TableModelChange"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[40] = type_
	type_ = &Type{
		"QAccessible2::TableModelChange*",
		classes["QAccessible2::TableModelChange"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[41] = type_
	type_ = &Type{
		"QAccessible2::TableModelChangeType",
		classes["QAccessible2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[42] = type_
	type_ = &Type{
		"QAccessible2Interface*",
		classes["QAccessible2Interface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[43] = type_
	type_ = &Type{
		"QAccessible::Action",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[44] = type_
	type_ = &Type{
		"QAccessible::Event",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[45] = type_
	type_ = &Type{
		"QAccessible::Method",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[46] = type_
	type_ = &Type{
		"QAccessible::RelationFlag",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[47] = type_
	type_ = &Type{
		"QAccessible::Role",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[48] = type_
	type_ = &Type{
		"QAccessible::StateFlag",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[49] = type_
	type_ = &Type{
		"QAccessible::Text",
		classes["QAccessible"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[50] = type_
	type_ = &Type{
		"QAccessibleActionInterface*",
		classes["QAccessibleActionInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[51] = type_
	type_ = &Type{
		"QAccessibleApplication*",
		classes["QAccessibleApplication"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[52] = type_
	type_ = &Type{
		"QAccessibleBridge*",
		classes["QAccessibleBridge"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[53] = type_
	type_ = &Type{
		"QAccessibleBridgeFactoryInterface*",
		classes["QAccessibleBridgeFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[54] = type_
	type_ = &Type{
		"QAccessibleBridgePlugin*",
		classes["QAccessibleBridgePlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[55] = type_
	type_ = &Type{
		"QAccessibleEditableTextInterface*",
		classes["QAccessibleEditableTextInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[56] = type_
	type_ = &Type{
		"QAccessibleEvent*",
		classes["QAccessibleEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[57] = type_
	type_ = &Type{
		"QAccessibleFactoryInterface*",
		classes["QAccessibleFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[58] = type_
	type_ = &Type{
		"QAccessibleImageInterface*",
		classes["QAccessibleImageInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[59] = type_
	type_ = &Type{
		"QAccessibleInterface*",
		classes["QAccessibleInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[60] = type_
	type_ = &Type{
		"QAccessibleInterface*(*)(const QString&,QObject*)",
		classes["QAccessibleInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[61] = type_
	type_ = &Type{
		"QAccessibleInterface**",
		classes["QAccessibleInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[62] = type_
	type_ = &Type{
		"QAccessibleInterfaceEx*",
		classes["QAccessibleInterfaceEx"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[63] = type_
	type_ = &Type{
		"QAccessibleObject*",
		classes["QAccessibleObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[64] = type_
	type_ = &Type{
		"QAccessibleObjectEx*",
		classes["QAccessibleObjectEx"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[65] = type_
	type_ = &Type{
		"QAccessiblePlugin*",
		classes["QAccessiblePlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[66] = type_
	type_ = &Type{
		"QAccessibleSimpleEditableTextInterface*",
		classes["QAccessibleSimpleEditableTextInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[67] = type_
	type_ = &Type{
		"QAccessibleTable2CellInterface*",
		classes["QAccessibleTable2CellInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[68] = type_
	type_ = &Type{
		"QAccessibleTable2Interface*",
		classes["QAccessibleTable2Interface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[69] = type_
	type_ = &Type{
		"QAccessibleTableInterface*",
		classes["QAccessibleTableInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[70] = type_
	type_ = &Type{
		"QAccessibleTextInterface*",
		classes["QAccessibleTextInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[71] = type_
	type_ = &Type{
		"QAccessibleValueInterface*",
		classes["QAccessibleValueInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[72] = type_
	type_ = &Type{
		"QAccessibleWidget*",
		classes["QAccessibleWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[73] = type_
	type_ = &Type{
		"QAccessibleWidgetEx*",
		classes["QAccessibleWidgetEx"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[74] = type_
	type_ = &Type{
		"QAction*",
		classes["QAction"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[75] = type_
	type_ = &Type{
		"QAction::ActionEvent",
		classes["QAction"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[76] = type_
	type_ = &Type{
		"QAction::MenuRole",
		classes["QAction"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[77] = type_
	type_ = &Type{
		"QAction::Priority",
		classes["QAction"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[78] = type_
	type_ = &Type{
		"QAction::SoftKeyRole",
		classes["QAction"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[79] = type_
	type_ = &Type{
		"QActionEvent*",
		classes["QActionEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[80] = type_
	type_ = &Type{
		"QActionGroup*",
		classes["QActionGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[81] = type_
	type_ = &Type{
		"QApplication*",
		classes["QApplication"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[82] = type_
	type_ = &Type{
		"QApplication::ColorSpec",
		classes["QApplication"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[83] = type_
	type_ = &Type{
		"QApplication::Type",
		classes["QApplication"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[84] = type_
	type_ = &Type{
		"QBitArray",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[85] = type_
	type_ = &Type{
		"QBitArray&",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[86] = type_
	type_ = &Type{
		"QBitmap",
		classes["QBitmap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[87] = type_
	type_ = &Type{
		"QBitmap&",
		classes["QBitmap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[88] = type_
	type_ = &Type{
		"QBitmap*",
		classes["QBitmap"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[89] = type_
	type_ = &Type{
		"QBool",
		classes["QBool"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[90] = type_
	type_ = &Type{
		"QBoxLayout*",
		classes["QBoxLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[91] = type_
	type_ = &Type{
		"QBoxLayout::Direction",
		classes["QBoxLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[92] = type_
	type_ = &Type{
		"QBrush",
		classes["QBrush"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[93] = type_
	type_ = &Type{
		"QBrush&",
		classes["QBrush"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[94] = type_
	type_ = &Type{
		"QBrush*",
		classes["QBrush"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[95] = type_
	type_ = &Type{
		"QButtonGroup*",
		classes["QButtonGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[96] = type_
	type_ = &Type{
		"QByteArray",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[97] = type_
	type_ = &Type{
		"QByteArray&",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[98] = type_
	type_ = &Type{
		"QCalendarWidget*",
		classes["QCalendarWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[99] = type_
	type_ = &Type{
		"QCalendarWidget::HorizontalHeaderFormat",
		classes["QCalendarWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[100] = type_
	type_ = &Type{
		"QCalendarWidget::SelectionMode",
		classes["QCalendarWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[101] = type_
	type_ = &Type{
		"QCalendarWidget::VerticalHeaderFormat",
		classes["QCalendarWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[102] = type_
	type_ = &Type{
		"QChar",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[103] = type_
	type_ = &Type{
		"QChar&",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[104] = type_
	type_ = &Type{
		"QCheckBox*",
		classes["QCheckBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[105] = type_
	type_ = &Type{
		"QChildEvent*",
		classes["QChildEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[106] = type_
	type_ = &Type{
		"QClipboard*",
		classes["QClipboard"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[107] = type_
	type_ = &Type{
		"QClipboard::Mode",
		classes["QClipboard"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[108] = type_
	type_ = &Type{
		"QClipboardEvent*",
		classes["QClipboardEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[109] = type_
	type_ = &Type{
		"QCloseEvent*",
		classes["QCloseEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[110] = type_
	type_ = &Type{
		"QColor",
		classes["QColor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[111] = type_
	type_ = &Type{
		"QColor&",
		classes["QColor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[112] = type_
	type_ = &Type{
		"QColor*",
		classes["QColor"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[113] = type_
	type_ = &Type{
		"QColor::Spec",
		classes["QColor"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[114] = type_
	type_ = &Type{
		"QColorDialog*",
		classes["QColorDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[115] = type_
	type_ = &Type{
		"QColorDialog::ColorDialogOption",
		classes["QColorDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[116] = type_
	type_ = &Type{
		"QColormap",
		classes["QColormap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[117] = type_
	type_ = &Type{
		"QColormap&",
		classes["QColormap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[118] = type_
	type_ = &Type{
		"QColormap*",
		classes["QColormap"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[119] = type_
	type_ = &Type{
		"QColormap::Mode",
		classes["QColormap"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[120] = type_
	type_ = &Type{
		"QColumnView*",
		classes["QColumnView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[121] = type_
	type_ = &Type{
		"QComboBox*",
		classes["QComboBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[122] = type_
	type_ = &Type{
		"QComboBox::InsertPolicy",
		classes["QComboBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[123] = type_
	type_ = &Type{
		"QComboBox::SizeAdjustPolicy",
		classes["QComboBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[124] = type_
	type_ = &Type{
		"QCommandLinkButton*",
		classes["QCommandLinkButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[125] = type_
	type_ = &Type{
		"QCommonStyle*",
		classes["QCommonStyle"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[126] = type_
	type_ = &Type{
		"QCompleter*",
		classes["QCompleter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[127] = type_
	type_ = &Type{
		"QCompleter::CompletionMode",
		classes["QCompleter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[128] = type_
	type_ = &Type{
		"QCompleter::ModelSorting",
		classes["QCompleter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[129] = type_
	type_ = &Type{
		"QConicalGradient*",
		classes["QConicalGradient"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[130] = type_
	type_ = &Type{
		"QContextMenuEvent*",
		classes["QContextMenuEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[131] = type_
	type_ = &Type{
		"QContextMenuEvent::Reason",
		classes["QContextMenuEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[132] = type_
	type_ = &Type{
		"QCursor",
		classes["QCursor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[133] = type_
	type_ = &Type{
		"QCursor&",
		classes["QCursor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[134] = type_
	type_ = &Type{
		"QCursor*",
		classes["QCursor"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[135] = type_
	type_ = &Type{
		"QDataStream&",
		classes["QDataStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[136] = type_
	type_ = &Type{
		"QDataWidgetMapper*",
		classes["QDataWidgetMapper"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[137] = type_
	type_ = &Type{
		"QDataWidgetMapper::SubmitPolicy",
		classes["QDataWidgetMapper"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[138] = type_
	type_ = &Type{
		"QDate",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[139] = type_
	type_ = &Type{
		"QDate&",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[140] = type_
	type_ = &Type{
		"QDateEdit*",
		classes["QDateEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[141] = type_
	type_ = &Type{
		"QDateTime",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[142] = type_
	type_ = &Type{
		"QDateTime&",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[143] = type_
	type_ = &Type{
		"QDateTimeEdit*",
		classes["QDateTimeEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[144] = type_
	type_ = &Type{
		"QDateTimeEdit::Section",
		classes["QDateTimeEdit"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[145] = type_
	type_ = &Type{
		"QDebug",
		classes["QDebug"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[146] = type_
	type_ = &Type{
		"QDesktopServices*",
		classes["QDesktopServices"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[147] = type_
	type_ = &Type{
		"QDesktopServices::StandardLocation",
		classes["QDesktopServices"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[148] = type_
	type_ = &Type{
		"QDesktopWidget*",
		classes["QDesktopWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[149] = type_
	type_ = &Type{
		"QDial*",
		classes["QDial"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[150] = type_
	type_ = &Type{
		"QDialog*",
		classes["QDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[151] = type_
	type_ = &Type{
		"QDialog::DialogCode",
		classes["QDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[152] = type_
	type_ = &Type{
		"QDialogButtonBox*",
		classes["QDialogButtonBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[153] = type_
	type_ = &Type{
		"QDialogButtonBox::ButtonLayout",
		classes["QDialogButtonBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[154] = type_
	type_ = &Type{
		"QDialogButtonBox::ButtonRole",
		classes["QDialogButtonBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[155] = type_
	type_ = &Type{
		"QDialogButtonBox::StandardButton",
		classes["QDialogButtonBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[156] = type_
	type_ = &Type{
		"QDir",
		classes["QDir"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[157] = type_
	type_ = &Type{
		"QDir::Filter",
		classes["QDir"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[158] = type_
	type_ = &Type{
		"QDir::SortFlag",
		classes["QDir"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[159] = type_
	type_ = &Type{
		"QDirIterator::IteratorFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[160] = type_
	type_ = &Type{
		"QDirModel*",
		classes["QDirModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[161] = type_
	type_ = &Type{
		"QDirModel::Roles",
		classes["QDirModel"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[162] = type_
	type_ = &Type{
		"QDockWidget*",
		classes["QDockWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[163] = type_
	type_ = &Type{
		"QDockWidget::DockWidgetFeature",
		classes["QDockWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[164] = type_
	type_ = &Type{
		"QDockWidget::DockWidgetFeatures",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[165] = type_
	type_ = &Type{
		"QDoubleSpinBox*",
		classes["QDoubleSpinBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[166] = type_
	type_ = &Type{
		"QDoubleValidator*",
		classes["QDoubleValidator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[167] = type_
	type_ = &Type{
		"QDoubleValidator::Notation",
		classes["QDoubleValidator"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[168] = type_
	type_ = &Type{
		"QDrag*",
		classes["QDrag"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[169] = type_
	type_ = &Type{
		"QDragEnterEvent*",
		classes["QDragEnterEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[170] = type_
	type_ = &Type{
		"QDragLeaveEvent*",
		classes["QDragLeaveEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[171] = type_
	type_ = &Type{
		"QDragMoveEvent*",
		classes["QDragMoveEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[172] = type_
	type_ = &Type{
		"QDragResponseEvent*",
		classes["QDragResponseEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[173] = type_
	type_ = &Type{
		"QDrawBorderPixmap::DrawingHint",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[174] = type_
	type_ = &Type{
		"QDropEvent*",
		classes["QDropEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[175] = type_
	type_ = &Type{
		"QEasingCurve&",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[176] = type_
	type_ = &Type{
		"QErrorMessage*",
		classes["QErrorMessage"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[177] = type_
	type_ = &Type{
		"QEvent*",
		classes["QEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[178] = type_
	type_ = &Type{
		"QEvent::Type",
		classes["QEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[179] = type_
	type_ = &Type{
		"QEventLoop::ProcessEventsFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[180] = type_
	type_ = &Type{
		"QEventPrivate*",
		classes["QEventPrivate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[181] = type_
	type_ = &Type{
		"QExplicitlySharedDataPointer<QPicturePrivate>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[182] = type_
	type_ = &Type{
		"QFile&",
		classes["QFile"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[183] = type_
	type_ = &Type{
		"QFile::Permission",
		classes["QFile"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[184] = type_
	type_ = &Type{
		"QFileDialog*",
		classes["QFileDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[185] = type_
	type_ = &Type{
		"QFileDialog::AcceptMode",
		classes["QFileDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[186] = type_
	type_ = &Type{
		"QFileDialog::DialogLabel",
		classes["QFileDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[187] = type_
	type_ = &Type{
		"QFileDialog::FileMode",
		classes["QFileDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[188] = type_
	type_ = &Type{
		"QFileDialog::Option",
		classes["QFileDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[189] = type_
	type_ = &Type{
		"QFileDialog::ViewMode",
		classes["QFileDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[190] = type_
	type_ = &Type{
		"QFileIconProvider*",
		classes["QFileIconProvider"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[191] = type_
	type_ = &Type{
		"QFileIconProvider::IconType",
		classes["QFileIconProvider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[192] = type_
	type_ = &Type{
		"QFileInfo",
		classes["QFileInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[193] = type_
	type_ = &Type{
		"QFileOpenEvent*",
		classes["QFileOpenEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[194] = type_
	type_ = &Type{
		"QFileSystemModel*",
		classes["QFileSystemModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[195] = type_
	type_ = &Type{
		"QFileSystemModel::Roles",
		classes["QFileSystemModel"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[196] = type_
	type_ = &Type{
		"QFlags<QAbstractFileEngine::FileFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[197] = type_
	type_ = &Type{
		"QFlags<QAbstractItemView::EditTrigger>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[198] = type_
	type_ = &Type{
		"QFlags<QAbstractPrintDialog::PrintDialogOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[199] = type_
	type_ = &Type{
		"QFlags<QAbstractSpinBox::StepEnabledFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[200] = type_
	type_ = &Type{
		"QFlags<QAccessible::RelationFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[201] = type_
	type_ = &Type{
		"QFlags<QAccessible::StateFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[202] = type_
	type_ = &Type{
		"QFlags<QColorDialog::ColorDialogOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[203] = type_
	type_ = &Type{
		"QFlags<QDateTimeEdit::Section>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[204] = type_
	type_ = &Type{
		"QFlags<QDialogButtonBox::StandardButton>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[205] = type_
	type_ = &Type{
		"QFlags<QDir::Filter>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[206] = type_
	type_ = &Type{
		"QFlags<QDir::SortFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[207] = type_
	type_ = &Type{
		"QFlags<QDirIterator::IteratorFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[208] = type_
	type_ = &Type{
		"QFlags<QDockWidget::DockWidgetFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[209] = type_
	type_ = &Type{
		"QFlags<QDrawBorderPixmap::DrawingHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[210] = type_
	type_ = &Type{
		"QFlags<QEventLoop::ProcessEventsFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[211] = type_
	type_ = &Type{
		"QFlags<QFile::Permission>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[212] = type_
	type_ = &Type{
		"QFlags<QFileDialog::Option>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[213] = type_
	type_ = &Type{
		"QFlags<QFontComboBox::FontFilter>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[214] = type_
	type_ = &Type{
		"QFlags<QFontDialog::FontDialogOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[215] = type_
	type_ = &Type{
		"QFlags<QGestureRecognizer::ResultFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[216] = type_
	type_ = &Type{
		"QFlags<QGraphicsBlurEffect::BlurHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[217] = type_
	type_ = &Type{
		"QFlags<QGraphicsEffect::ChangeFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[218] = type_
	type_ = &Type{
		"QFlags<QGraphicsItem::GraphicsItemFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[219] = type_
	type_ = &Type{
		"QFlags<QGraphicsScene::SceneLayer>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[220] = type_
	type_ = &Type{
		"QFlags<QGraphicsView::CacheModeFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[221] = type_
	type_ = &Type{
		"QFlags<QGraphicsView::OptimizationFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[222] = type_
	type_ = &Type{
		"QFlags<QIODevice::OpenModeFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[223] = type_
	type_ = &Type{
		"QFlags<QImageIOPlugin::Capability>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[224] = type_
	type_ = &Type{
		"QFlags<QInputDialog::InputDialogOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[225] = type_
	type_ = &Type{
		"QFlags<QItemSelectionModel::SelectionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[226] = type_
	type_ = &Type{
		"QFlags<QLibrary::LoadHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[227] = type_
	type_ = &Type{
		"QFlags<QLocale::NumberOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[228] = type_
	type_ = &Type{
		"QFlags<QMainWindow::DockOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[229] = type_
	type_ = &Type{
		"QFlags<QMdiArea::AreaOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[230] = type_
	type_ = &Type{
		"QFlags<QMdiSubWindow::SubWindowOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[231] = type_
	type_ = &Type{
		"QFlags<QMessageBox::StandardButton>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[232] = type_
	type_ = &Type{
		"QFlags<QPageSetupDialog::PageSetupDialogOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[233] = type_
	type_ = &Type{
		"QFlags<QPaintEngine::DirtyFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[234] = type_
	type_ = &Type{
		"QFlags<QPaintEngine::PaintEngineFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[235] = type_
	type_ = &Type{
		"QFlags<QPainter::PixmapFragmentHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[236] = type_
	type_ = &Type{
		"QFlags<QPainter::RenderHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[237] = type_
	type_ = &Type{
		"QFlags<QPinchGesture::ChangeFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[238] = type_
	type_ = &Type{
		"QFlags<QSizePolicy::ControlType>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[239] = type_
	type_ = &Type{
		"QFlags<QString::SectionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[240] = type_
	type_ = &Type{
		"QFlags<QStyle::StateFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[241] = type_
	type_ = &Type{
		"QFlags<QStyle::SubControl>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[242] = type_
	type_ = &Type{
		"QFlags<QStyleOptionButton::ButtonFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[243] = type_
	type_ = &Type{
		"QFlags<QStyleOptionFrameV2::FrameFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[244] = type_
	type_ = &Type{
		"QFlags<QStyleOptionQ3ListViewItem::Q3ListViewItemFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[245] = type_
	type_ = &Type{
		"QFlags<QStyleOptionTab::CornerWidget>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[246] = type_
	type_ = &Type{
		"QFlags<QStyleOptionToolBar::ToolBarFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[247] = type_
	type_ = &Type{
		"QFlags<QStyleOptionToolButton::ToolButtonFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[248] = type_
	type_ = &Type{
		"QFlags<QStyleOptionViewItemV2::ViewItemFeature>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[249] = type_
	type_ = &Type{
		"QFlags<QTextCodec::ConversionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[250] = type_
	type_ = &Type{
		"QFlags<QTextDocument::FindFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[251] = type_
	type_ = &Type{
		"QFlags<QTextEdit::AutoFormattingFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[252] = type_
	type_ = &Type{
		"QFlags<QTextFormat::PageBreakFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[253] = type_
	type_ = &Type{
		"QFlags<QTextItem::RenderFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[254] = type_
	type_ = &Type{
		"QFlags<QTextOption::Flag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[255] = type_
	type_ = &Type{
		"QFlags<QTextStream::NumberFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[256] = type_
	type_ = &Type{
		"QFlags<QTreeWidgetItemIterator::IteratorFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[257] = type_
	type_ = &Type{
		"QFlags<QUrl::FormattingOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[258] = type_
	type_ = &Type{
		"QFlags<QWidget::RenderFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[259] = type_
	type_ = &Type{
		"QFlags<QWizard::WizardOption>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[260] = type_
	type_ = &Type{
		"QFlags<Qt::AlignmentFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[261] = type_
	type_ = &Type{
		"QFlags<Qt::DockWidgetArea>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[262] = type_
	type_ = &Type{
		"QFlags<Qt::DropAction>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[263] = type_
	type_ = &Type{
		"QFlags<Qt::GestureFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[264] = type_
	type_ = &Type{
		"QFlags<Qt::ImageConversionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[265] = type_
	type_ = &Type{
		"QFlags<Qt::InputMethodHint>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[266] = type_
	type_ = &Type{
		"QFlags<Qt::ItemFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[267] = type_
	type_ = &Type{
		"QFlags<Qt::KeyboardModifier>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[268] = type_
	type_ = &Type{
		"QFlags<Qt::MatchFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[269] = type_
	type_ = &Type{
		"QFlags<Qt::MouseButton>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[270] = type_
	type_ = &Type{
		"QFlags<Qt::Orientation>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[271] = type_
	type_ = &Type{
		"QFlags<Qt::TextInteractionFlag>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[272] = type_
	type_ = &Type{
		"QFlags<Qt::ToolBarArea>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[273] = type_
	type_ = &Type{
		"QFlags<Qt::TouchPointState>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[274] = type_
	type_ = &Type{
		"QFlags<Qt::WindowState>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[275] = type_
	type_ = &Type{
		"QFlags<Qt::WindowType>",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[276] = type_
	type_ = &Type{
		"QFocusEvent*",
		classes["QFocusEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[277] = type_
	type_ = &Type{
		"QFocusFrame*",
		classes["QFocusFrame"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[278] = type_
	type_ = &Type{
		"QFont",
		classes["QFont"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[279] = type_
	type_ = &Type{
		"QFont&",
		classes["QFont"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[280] = type_
	type_ = &Type{
		"QFont*",
		classes["QFont"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[281] = type_
	type_ = &Type{
		"QFont::Capitalization",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[282] = type_
	type_ = &Type{
		"QFont::HintingPreference",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[283] = type_
	type_ = &Type{
		"QFont::ResolveProperties",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[284] = type_
	type_ = &Type{
		"QFont::SpacingType",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[285] = type_
	type_ = &Type{
		"QFont::Stretch",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[286] = type_
	type_ = &Type{
		"QFont::Style",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[287] = type_
	type_ = &Type{
		"QFont::StyleHint",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[288] = type_
	type_ = &Type{
		"QFont::StyleStrategy",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[289] = type_
	type_ = &Type{
		"QFont::Weight",
		classes["QFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[290] = type_
	type_ = &Type{
		"QFontComboBox*",
		classes["QFontComboBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[291] = type_
	type_ = &Type{
		"QFontComboBox::FontFilter",
		classes["QFontComboBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[292] = type_
	type_ = &Type{
		"QFontDatabase*",
		classes["QFontDatabase"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[293] = type_
	type_ = &Type{
		"QFontDatabase::WritingSystem",
		classes["QFontDatabase"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[294] = type_
	type_ = &Type{
		"QFontDialog*",
		classes["QFontDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[295] = type_
	type_ = &Type{
		"QFontDialog::FontDialogOption",
		classes["QFontDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[296] = type_
	type_ = &Type{
		"QFontInfo",
		classes["QFontInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[297] = type_
	type_ = &Type{
		"QFontInfo&",
		classes["QFontInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[298] = type_
	type_ = &Type{
		"QFontInfo*",
		classes["QFontInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[299] = type_
	type_ = &Type{
		"QFontMetrics",
		classes["QFontMetrics"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[300] = type_
	type_ = &Type{
		"QFontMetrics&",
		classes["QFontMetrics"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[301] = type_
	type_ = &Type{
		"QFontMetrics*",
		classes["QFontMetrics"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[302] = type_
	type_ = &Type{
		"QFontMetricsF&",
		classes["QFontMetricsF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[303] = type_
	type_ = &Type{
		"QFontMetricsF*",
		classes["QFontMetricsF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[304] = type_
	type_ = &Type{
		"QFormLayout*",
		classes["QFormLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[305] = type_
	type_ = &Type{
		"QFormLayout::FieldGrowthPolicy",
		classes["QFormLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[306] = type_
	type_ = &Type{
		"QFormLayout::ItemRole",
		classes["QFormLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[307] = type_
	type_ = &Type{
		"QFormLayout::ItemRole*",
		classes["QFormLayout"],
		T_ENUM,
		KIND_POINTER,
		false,
	}
	types[308] = type_
	type_ = &Type{
		"QFormLayout::RowWrapPolicy",
		classes["QFormLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[309] = type_
	type_ = &Type{
		"QFrame*",
		classes["QFrame"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[310] = type_
	type_ = &Type{
		"QFrame::Shadow",
		classes["QFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[311] = type_
	type_ = &Type{
		"QFrame::Shape",
		classes["QFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[312] = type_
	type_ = &Type{
		"QFrame::StyleMask",
		classes["QFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[313] = type_
	type_ = &Type{
		"QGenericMatrix<3,3,double>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[314] = type_
	type_ = &Type{
		"QGesture*",
		classes["QGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[315] = type_
	type_ = &Type{
		"QGesture::GestureCancelPolicy",
		classes["QGesture"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[316] = type_
	type_ = &Type{
		"QGestureEvent*",
		classes["QGestureEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[317] = type_
	type_ = &Type{
		"QGestureRecognizer*",
		classes["QGestureRecognizer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[318] = type_
	type_ = &Type{
		"QGestureRecognizer::ResultFlag",
		classes["QGestureRecognizer"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[319] = type_
	type_ = &Type{
		"QGlyphRun&",
		classes["QGlyphRun"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[320] = type_
	type_ = &Type{
		"QGlyphRun*",
		classes["QGlyphRun"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[321] = type_
	type_ = &Type{
		"QGradient*",
		classes["QGradient"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[322] = type_
	type_ = &Type{
		"QGradient::CoordinateMode",
		classes["QGradient"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[323] = type_
	type_ = &Type{
		"QGradient::InterpolationMode",
		classes["QGradient"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[324] = type_
	type_ = &Type{
		"QGradient::Spread",
		classes["QGradient"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[325] = type_
	type_ = &Type{
		"QGradient::Type",
		classes["QGradient"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[326] = type_
	type_ = &Type{
		"QGraphicsAnchor*",
		classes["QGraphicsAnchor"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[327] = type_
	type_ = &Type{
		"QGraphicsAnchorLayout*",
		classes["QGraphicsAnchorLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[328] = type_
	type_ = &Type{
		"QGraphicsBlurEffect*",
		classes["QGraphicsBlurEffect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[329] = type_
	type_ = &Type{
		"QGraphicsBlurEffect::BlurHint",
		classes["QGraphicsBlurEffect"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[330] = type_
	type_ = &Type{
		"QGraphicsBlurEffect::BlurHints",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[331] = type_
	type_ = &Type{
		"QGraphicsColorizeEffect*",
		classes["QGraphicsColorizeEffect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[332] = type_
	type_ = &Type{
		"QGraphicsDropShadowEffect*",
		classes["QGraphicsDropShadowEffect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[333] = type_
	type_ = &Type{
		"QGraphicsEffect*",
		classes["QGraphicsEffect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[334] = type_
	type_ = &Type{
		"QGraphicsEffect::ChangeFlag",
		classes["QGraphicsEffect"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[335] = type_
	type_ = &Type{
		"QGraphicsEffect::PixmapPadMode",
		classes["QGraphicsEffect"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[336] = type_
	type_ = &Type{
		"QGraphicsEffectSource*",
		classes["QGraphicsEffectSource"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[337] = type_
	type_ = &Type{
		"QGraphicsEllipseItem*",
		classes["QGraphicsEllipseItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[338] = type_
	type_ = &Type{
		"QGraphicsGridLayout*",
		classes["QGraphicsGridLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[339] = type_
	type_ = &Type{
		"QGraphicsItem*",
		classes["QGraphicsItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[340] = type_
	type_ = &Type{
		"QGraphicsItem**",
		classes["QGraphicsItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[341] = type_
	type_ = &Type{
		"QGraphicsItem::CacheMode",
		classes["QGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[342] = type_
	type_ = &Type{
		"QGraphicsItem::Extension",
		classes["QGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[343] = type_
	type_ = &Type{
		"QGraphicsItem::GraphicsItemChange",
		classes["QGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[344] = type_
	type_ = &Type{
		"QGraphicsItem::GraphicsItemFlag",
		classes["QGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[345] = type_
	type_ = &Type{
		"QGraphicsItem::PanelModality",
		classes["QGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[346] = type_
	type_ = &Type{
		"QGraphicsItemAnimation*",
		classes["QGraphicsItemAnimation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[347] = type_
	type_ = &Type{
		"QGraphicsItemGroup*",
		classes["QGraphicsItemGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[348] = type_
	type_ = &Type{
		"QGraphicsLayout*",
		classes["QGraphicsLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[349] = type_
	type_ = &Type{
		"QGraphicsLayoutItem*",
		classes["QGraphicsLayoutItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[350] = type_
	type_ = &Type{
		"QGraphicsLineItem*",
		classes["QGraphicsLineItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[351] = type_
	type_ = &Type{
		"QGraphicsLinearLayout*",
		classes["QGraphicsLinearLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[352] = type_
	type_ = &Type{
		"QGraphicsObject*",
		classes["QGraphicsObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[353] = type_
	type_ = &Type{
		"QGraphicsOpacityEffect*",
		classes["QGraphicsOpacityEffect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[354] = type_
	type_ = &Type{
		"QGraphicsPathItem*",
		classes["QGraphicsPathItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[355] = type_
	type_ = &Type{
		"QGraphicsPixmapItem*",
		classes["QGraphicsPixmapItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[356] = type_
	type_ = &Type{
		"QGraphicsPixmapItem::ShapeMode",
		classes["QGraphicsPixmapItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[357] = type_
	type_ = &Type{
		"QGraphicsPolygonItem*",
		classes["QGraphicsPolygonItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[358] = type_
	type_ = &Type{
		"QGraphicsProxyWidget*",
		classes["QGraphicsProxyWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[359] = type_
	type_ = &Type{
		"QGraphicsRectItem*",
		classes["QGraphicsRectItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[360] = type_
	type_ = &Type{
		"QGraphicsRotation*",
		classes["QGraphicsRotation"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[361] = type_
	type_ = &Type{
		"QGraphicsScale*",
		classes["QGraphicsScale"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[362] = type_
	type_ = &Type{
		"QGraphicsScene*",
		classes["QGraphicsScene"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[363] = type_
	type_ = &Type{
		"QGraphicsScene::ItemIndexMethod",
		classes["QGraphicsScene"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[364] = type_
	type_ = &Type{
		"QGraphicsScene::SceneLayer",
		classes["QGraphicsScene"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[365] = type_
	type_ = &Type{
		"QGraphicsScene::SceneLayers",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[366] = type_
	type_ = &Type{
		"QGraphicsSceneContextMenuEvent*",
		classes["QGraphicsSceneContextMenuEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[367] = type_
	type_ = &Type{
		"QGraphicsSceneContextMenuEvent::Reason",
		classes["QGraphicsSceneContextMenuEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[368] = type_
	type_ = &Type{
		"QGraphicsSceneDragDropEvent*",
		classes["QGraphicsSceneDragDropEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[369] = type_
	type_ = &Type{
		"QGraphicsSceneEvent*",
		classes["QGraphicsSceneEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[370] = type_
	type_ = &Type{
		"QGraphicsSceneEventPrivate*",
		classes["QGraphicsSceneEventPrivate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[371] = type_
	type_ = &Type{
		"QGraphicsSceneHelpEvent*",
		classes["QGraphicsSceneHelpEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[372] = type_
	type_ = &Type{
		"QGraphicsSceneHoverEvent*",
		classes["QGraphicsSceneHoverEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[373] = type_
	type_ = &Type{
		"QGraphicsSceneMouseEvent*",
		classes["QGraphicsSceneMouseEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[374] = type_
	type_ = &Type{
		"QGraphicsSceneMoveEvent*",
		classes["QGraphicsSceneMoveEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[375] = type_
	type_ = &Type{
		"QGraphicsSceneResizeEvent*",
		classes["QGraphicsSceneResizeEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[376] = type_
	type_ = &Type{
		"QGraphicsSceneWheelEvent*",
		classes["QGraphicsSceneWheelEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[377] = type_
	type_ = &Type{
		"QGraphicsSimpleTextItem*",
		classes["QGraphicsSimpleTextItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[378] = type_
	type_ = &Type{
		"QGraphicsTextItem*",
		classes["QGraphicsTextItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[379] = type_
	type_ = &Type{
		"QGraphicsTransform*",
		classes["QGraphicsTransform"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[380] = type_
	type_ = &Type{
		"QGraphicsView*",
		classes["QGraphicsView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[381] = type_
	type_ = &Type{
		"QGraphicsView::CacheModeFlag",
		classes["QGraphicsView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[382] = type_
	type_ = &Type{
		"QGraphicsView::DragMode",
		classes["QGraphicsView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[383] = type_
	type_ = &Type{
		"QGraphicsView::OptimizationFlag",
		classes["QGraphicsView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[384] = type_
	type_ = &Type{
		"QGraphicsView::ViewportAnchor",
		classes["QGraphicsView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[385] = type_
	type_ = &Type{
		"QGraphicsView::ViewportUpdateMode",
		classes["QGraphicsView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[386] = type_
	type_ = &Type{
		"QGraphicsWidget*",
		classes["QGraphicsWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[387] = type_
	type_ = &Type{
		"QGridLayout*",
		classes["QGridLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[388] = type_
	type_ = &Type{
		"QGroupBox*",
		classes["QGroupBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[389] = type_
	type_ = &Type{
		"QHBoxLayout*",
		classes["QHBoxLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[390] = type_
	type_ = &Type{
		"QHeaderView*",
		classes["QHeaderView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[391] = type_
	type_ = &Type{
		"QHeaderView::ResizeMode",
		classes["QHeaderView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[392] = type_
	type_ = &Type{
		"QHelpEvent*",
		classes["QHelpEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[393] = type_
	type_ = &Type{
		"QHideEvent*",
		classes["QHideEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[394] = type_
	type_ = &Type{
		"QHoverEvent*",
		classes["QHoverEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[395] = type_
	type_ = &Type{
		"QIODevice*",
		classes["QIODevice"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[396] = type_
	type_ = &Type{
		"QIODevice::OpenModeFlag",
		classes["QIODevice"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[397] = type_
	type_ = &Type{
		"QIcon",
		classes["QIcon"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[398] = type_
	type_ = &Type{
		"QIcon&",
		classes["QIcon"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[399] = type_
	type_ = &Type{
		"QIcon*",
		classes["QIcon"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[400] = type_
	type_ = &Type{
		"QIcon::Mode",
		classes["QIcon"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[401] = type_
	type_ = &Type{
		"QIcon::State",
		classes["QIcon"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[402] = type_
	type_ = &Type{
		"QIconDragEvent*",
		classes["QIconDragEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[403] = type_
	type_ = &Type{
		"QIconEngine*",
		classes["QIconEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[404] = type_
	type_ = &Type{
		"QIconEngineFactoryInterface*",
		classes["QIconEngineFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[405] = type_
	type_ = &Type{
		"QIconEngineFactoryInterfaceV2*",
		classes["QIconEngineFactoryInterfaceV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[406] = type_
	type_ = &Type{
		"QIconEnginePlugin*",
		classes["QIconEnginePlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[407] = type_
	type_ = &Type{
		"QIconEnginePluginV2*",
		classes["QIconEnginePluginV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[408] = type_
	type_ = &Type{
		"QIconEngineV2*",
		classes["QIconEngineV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[409] = type_
	type_ = &Type{
		"QIconEngineV2::AvailableSizesArgument*",
		classes["QIconEngineV2::AvailableSizesArgument"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[410] = type_
	type_ = &Type{
		"QIconEngineV2::IconEngineHook",
		classes["QIconEngineV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[411] = type_
	type_ = &Type{
		"QIconPrivate*&",
		classes["QIconPrivate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[412] = type_
	type_ = &Type{
		"QImage",
		classes["QImage"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[413] = type_
	type_ = &Type{
		"QImage&",
		classes["QImage"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[414] = type_
	type_ = &Type{
		"QImage*",
		classes["QImage"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[415] = type_
	type_ = &Type{
		"QImage::Format",
		classes["QImage"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[416] = type_
	type_ = &Type{
		"QImage::InvertMode",
		classes["QImage"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[417] = type_
	type_ = &Type{
		"QImageData*&",
		classes["QImageData"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[418] = type_
	type_ = &Type{
		"QImageIOHandler*",
		classes["QImageIOHandler"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[419] = type_
	type_ = &Type{
		"QImageIOHandler::ImageOption",
		classes["QImageIOHandler"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[420] = type_
	type_ = &Type{
		"QImageIOHandlerFactoryInterface*",
		classes["QImageIOHandlerFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[421] = type_
	type_ = &Type{
		"QImageIOPlugin*",
		classes["QImageIOPlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[422] = type_
	type_ = &Type{
		"QImageIOPlugin::Capability",
		classes["QImageIOPlugin"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[423] = type_
	type_ = &Type{
		"QImageReader*",
		classes["QImageReader"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[424] = type_
	type_ = &Type{
		"QImageReader::ImageReaderError",
		classes["QImageReader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[425] = type_
	type_ = &Type{
		"QImageTextKeyLang*",
		classes["QImageTextKeyLang"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[426] = type_
	type_ = &Type{
		"QImageWriter*",
		classes["QImageWriter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[427] = type_
	type_ = &Type{
		"QImageWriter::ImageWriterError",
		classes["QImageWriter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[428] = type_
	type_ = &Type{
		"QIncompatibleFlag",
		classes["QIncompatibleFlag"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[429] = type_
	type_ = &Type{
		"QInputContext*",
		classes["QInputContext"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[430] = type_
	type_ = &Type{
		"QInputContext::StandardFormat",
		classes["QInputContext"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[431] = type_
	type_ = &Type{
		"QInputContextFactory*",
		classes["QInputContextFactory"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[432] = type_
	type_ = &Type{
		"QInputContextFactoryInterface*",
		classes["QInputContextFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[433] = type_
	type_ = &Type{
		"QInputContextPlugin*",
		classes["QInputContextPlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[434] = type_
	type_ = &Type{
		"QInputDialog*",
		classes["QInputDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[435] = type_
	type_ = &Type{
		"QInputDialog::InputDialogOption",
		classes["QInputDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[436] = type_
	type_ = &Type{
		"QInputDialog::InputMode",
		classes["QInputDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[437] = type_
	type_ = &Type{
		"QInputEvent*",
		classes["QInputEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[438] = type_
	type_ = &Type{
		"QInputMethodEvent*",
		classes["QInputMethodEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[439] = type_
	type_ = &Type{
		"QInputMethodEvent::Attribute*",
		classes["QInputMethodEvent::Attribute"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[440] = type_
	type_ = &Type{
		"QInputMethodEvent::AttributeType",
		classes["QInputMethodEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[441] = type_
	type_ = &Type{
		"QIntValidator*",
		classes["QIntValidator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[442] = type_
	type_ = &Type{
		"QItemDelegate*",
		classes["QItemDelegate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[443] = type_
	type_ = &Type{
		"QItemEditorCreatorBase*",
		classes["QItemEditorCreatorBase"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[444] = type_
	type_ = &Type{
		"QItemEditorFactory*",
		classes["QItemEditorFactory"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[445] = type_
	type_ = &Type{
		"QItemSelection",
		classes["QItemSelection"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[446] = type_
	type_ = &Type{
		"QItemSelection*",
		classes["QItemSelection"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[447] = type_
	type_ = &Type{
		"QItemSelectionModel*",
		classes["QItemSelectionModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[448] = type_
	type_ = &Type{
		"QItemSelectionModel::SelectionFlag",
		classes["QItemSelectionModel"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[449] = type_
	type_ = &Type{
		"QItemSelectionModel::SelectionFlags",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[450] = type_
	type_ = &Type{
		"QItemSelectionRange",
		classes["QItemSelectionRange"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[451] = type_
	type_ = &Type{
		"QItemSelectionRange*",
		classes["QItemSelectionRange"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[452] = type_
	type_ = &Type{
		"QKeyEvent*",
		classes["QKeyEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[453] = type_
	type_ = &Type{
		"QKeyEventTransition*",
		classes["QKeyEventTransition"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[454] = type_
	type_ = &Type{
		"QKeySequence",
		classes["QKeySequence"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[455] = type_
	type_ = &Type{
		"QKeySequence&",
		classes["QKeySequence"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[456] = type_
	type_ = &Type{
		"QKeySequence*",
		classes["QKeySequence"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[457] = type_
	type_ = &Type{
		"QKeySequence::SequenceFormat",
		classes["QKeySequence"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[458] = type_
	type_ = &Type{
		"QKeySequence::SequenceMatch",
		classes["QKeySequence"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[459] = type_
	type_ = &Type{
		"QKeySequence::StandardKey",
		classes["QKeySequence"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[460] = type_
	type_ = &Type{
		"QKeySequencePrivate*&",
		classes["QKeySequencePrivate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[461] = type_
	type_ = &Type{
		"QLCDNumber*",
		classes["QLCDNumber"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[462] = type_
	type_ = &Type{
		"QLCDNumber::Mode",
		classes["QLCDNumber"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[463] = type_
	type_ = &Type{
		"QLCDNumber::SegmentStyle",
		classes["QLCDNumber"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[464] = type_
	type_ = &Type{
		"QLabel*",
		classes["QLabel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[465] = type_
	type_ = &Type{
		"QLayout*",
		classes["QLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[466] = type_
	type_ = &Type{
		"QLayout::SizeConstraint",
		classes["QLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[467] = type_
	type_ = &Type{
		"QLayoutItem*",
		classes["QLayoutItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[468] = type_
	type_ = &Type{
		"QLibrary::LoadHint",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[469] = type_
	type_ = &Type{
		"QLine",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[470] = type_
	type_ = &Type{
		"QLine&",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[471] = type_
	type_ = &Type{
		"QLineEdit*",
		classes["QLineEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[472] = type_
	type_ = &Type{
		"QLineEdit::EchoMode",
		classes["QLineEdit"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[473] = type_
	type_ = &Type{
		"QLineF",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[474] = type_
	type_ = &Type{
		"QLineF&",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[475] = type_
	type_ = &Type{
		"QLinearGradient*",
		classes["QLinearGradient"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[476] = type_
	type_ = &Type{
		"QList<QAbstractButton*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[477] = type_
	type_ = &Type{
		"QList<QAccessibleInterface*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[478] = type_
	type_ = &Type{
		"QList<QAccessibleTable2CellInterface*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[479] = type_
	type_ = &Type{
		"QList<QAction*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[480] = type_
	type_ = &Type{
		"QList<QByteArray>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[481] = type_
	type_ = &Type{
		"QList<QByteArray>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[482] = type_
	type_ = &Type{
		"QList<QDockWidget*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[483] = type_
	type_ = &Type{
		"QList<QFontDatabase::WritingSystem>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[484] = type_
	type_ = &Type{
		"QList<QGesture*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[485] = type_
	type_ = &Type{
		"QList<QGlyphRun>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[486] = type_
	type_ = &Type{
		"QList<QGraphicsItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[487] = type_
	type_ = &Type{
		"QList<QGraphicsTransform*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[488] = type_
	type_ = &Type{
		"QList<QGraphicsView*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[489] = type_
	type_ = &Type{
		"QList<QGraphicsWidget*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[490] = type_
	type_ = &Type{
		"QList<QImageTextKeyLang>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[491] = type_
	type_ = &Type{
		"QList<QKeySequence>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[492] = type_
	type_ = &Type{
		"QList<QListWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[493] = type_
	type_ = &Type{
		"QList<QMdiSubWindow*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[494] = type_
	type_ = &Type{
		"QList<QModelIndex>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[495] = type_
	type_ = &Type{
		"QList<QPair<double,QPointF> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[496] = type_
	type_ = &Type{
		"QList<QPair<double,double> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[497] = type_
	type_ = &Type{
		"QList<QPolygonF>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[498] = type_
	type_ = &Type{
		"QList<QPrinter::PageSize>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[499] = type_
	type_ = &Type{
		"QList<QPrinterInfo>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[500] = type_
	type_ = &Type{
		"QList<QRectF>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[501] = type_
	type_ = &Type{
		"QList<QSize>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[502] = type_
	type_ = &Type{
		"QList<QSize>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[503] = type_
	type_ = &Type{
		"QList<QStandardItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[504] = type_
	type_ = &Type{
		"QList<QTableWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[505] = type_
	type_ = &Type{
		"QList<QTableWidgetSelectionRange>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[506] = type_
	type_ = &Type{
		"QList<QTextBlock>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[507] = type_
	type_ = &Type{
		"QList<QTextEdit::ExtraSelection>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[508] = type_
	type_ = &Type{
		"QList<QTextFrame*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[509] = type_
	type_ = &Type{
		"QList<QTextLayout::FormatRange>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[510] = type_
	type_ = &Type{
		"QList<QTextOption::Tab>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[511] = type_
	type_ = &Type{
		"QList<QTouchEvent::TouchPoint>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[512] = type_
	type_ = &Type{
		"QList<QTreeWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[513] = type_
	type_ = &Type{
		"QList<QUndoStack*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[514] = type_
	type_ = &Type{
		"QList<QUrl>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[515] = type_
	type_ = &Type{
		"QList<QWidget*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[516] = type_
	type_ = &Type{
		"QList<double>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[517] = type_
	type_ = &Type{
		"QList<int>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[518] = type_
	type_ = &Type{
		"QList<int>*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[519] = type_
	type_ = &Type{
		"QList<void*>*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[520] = type_
	type_ = &Type{
		"QListView*",
		classes["QListView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[521] = type_
	type_ = &Type{
		"QListView::Flow",
		classes["QListView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[522] = type_
	type_ = &Type{
		"QListView::LayoutMode",
		classes["QListView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[523] = type_
	type_ = &Type{
		"QListView::Movement",
		classes["QListView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[524] = type_
	type_ = &Type{
		"QListView::ResizeMode",
		classes["QListView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[525] = type_
	type_ = &Type{
		"QListView::ViewMode",
		classes["QListView"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[526] = type_
	type_ = &Type{
		"QListWidget*",
		classes["QListWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[527] = type_
	type_ = &Type{
		"QListWidgetItem&",
		classes["QListWidgetItem"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[528] = type_
	type_ = &Type{
		"QListWidgetItem*",
		classes["QListWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[529] = type_
	type_ = &Type{
		"QListWidgetItem::ItemType",
		classes["QListWidgetItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[530] = type_
	type_ = &Type{
		"QLocale",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[531] = type_
	type_ = &Type{
		"QLocale&",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[532] = type_
	type_ = &Type{
		"QLocale::NumberOption",
		classes["QLocale"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[533] = type_
	type_ = &Type{
		"QMainWindow*",
		classes["QMainWindow"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[534] = type_
	type_ = &Type{
		"QMainWindow::DockOption",
		classes["QMainWindow"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[535] = type_
	type_ = &Type{
		"QMap<QDate,QTextCharFormat>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[536] = type_
	type_ = &Type{
		"QMap<int,QVariant>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[537] = type_
	type_ = &Type{
		"QMargins",
		classes["QMargins"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[538] = type_
	type_ = &Type{
		"QMatrix",
		classes["QMatrix"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[539] = type_
	type_ = &Type{
		"QMatrix&",
		classes["QMatrix"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[540] = type_
	type_ = &Type{
		"QMatrix*",
		classes["QMatrix"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[541] = type_
	type_ = &Type{
		"QMatrix4x4",
		classes["QMatrix4x4"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[542] = type_
	type_ = &Type{
		"QMatrix4x4&",
		classes["QMatrix4x4"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[543] = type_
	type_ = &Type{
		"QMatrix4x4*",
		classes["QMatrix4x4"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[544] = type_
	type_ = &Type{
		"QMdiArea*",
		classes["QMdiArea"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[545] = type_
	type_ = &Type{
		"QMdiArea::AreaOption",
		classes["QMdiArea"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[546] = type_
	type_ = &Type{
		"QMdiArea::ViewMode",
		classes["QMdiArea"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[547] = type_
	type_ = &Type{
		"QMdiArea::WindowOrder",
		classes["QMdiArea"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[548] = type_
	type_ = &Type{
		"QMdiSubWindow*",
		classes["QMdiSubWindow"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[549] = type_
	type_ = &Type{
		"QMdiSubWindow::SubWindowOption",
		classes["QMdiSubWindow"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[550] = type_
	type_ = &Type{
		"QMenu*",
		classes["QMenu"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[551] = type_
	type_ = &Type{
		"QMenuBar*",
		classes["QMenuBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[552] = type_
	type_ = &Type{
		"QMessageBox*",
		classes["QMessageBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[553] = type_
	type_ = &Type{
		"QMessageBox::ButtonRole",
		classes["QMessageBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[554] = type_
	type_ = &Type{
		"QMessageBox::Icon",
		classes["QMessageBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[555] = type_
	type_ = &Type{
		"QMessageBox::StandardButton",
		classes["QMessageBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[556] = type_
	type_ = &Type{
		"QMetaObject::Call",
		classes["QMetaObject"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[557] = type_
	type_ = &Type{
		"QMimeData*",
		classes["QMimeData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[558] = type_
	type_ = &Type{
		"QMimeSource*",
		classes["QMimeSource"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[559] = type_
	type_ = &Type{
		"QModelIndex",
		classes["QModelIndex"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[560] = type_
	type_ = &Type{
		"QModelIndex&",
		classes["QModelIndex"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[561] = type_
	type_ = &Type{
		"QModelIndexList",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[562] = type_
	type_ = &Type{
		"QMouseEvent*",
		classes["QMouseEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[563] = type_
	type_ = &Type{
		"QMouseEventTransition*",
		classes["QMouseEventTransition"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[564] = type_
	type_ = &Type{
		"QMoveEvent*",
		classes["QMoveEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[565] = type_
	type_ = &Type{
		"QMovie*",
		classes["QMovie"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[566] = type_
	type_ = &Type{
		"QMovie::CacheMode",
		classes["QMovie"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[567] = type_
	type_ = &Type{
		"QMovie::MovieState",
		classes["QMovie"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[568] = type_
	type_ = &Type{
		"QObject*",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[569] = type_
	type_ = &Type{
		"QObject*(*)()",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[570] = type_
	type_ = &Type{
		"QPageSetupDialog*",
		classes["QPageSetupDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[571] = type_
	type_ = &Type{
		"QPageSetupDialog::PageSetupDialogOption",
		classes["QPageSetupDialog"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[572] = type_
	type_ = &Type{
		"QPaintDevice*",
		classes["QPaintDevice"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[573] = type_
	type_ = &Type{
		"QPaintDevice::PaintDeviceMetric",
		classes["QPaintDevice"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[574] = type_
	type_ = &Type{
		"QPaintEngine*",
		classes["QPaintEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[575] = type_
	type_ = &Type{
		"QPaintEngine::DirtyFlag",
		classes["QPaintEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[576] = type_
	type_ = &Type{
		"QPaintEngine::PaintEngineFeature",
		classes["QPaintEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[577] = type_
	type_ = &Type{
		"QPaintEngine::PolygonDrawMode",
		classes["QPaintEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[578] = type_
	type_ = &Type{
		"QPaintEngine::Type",
		classes["QPaintEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[579] = type_
	type_ = &Type{
		"QPaintEngineState*",
		classes["QPaintEngineState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[580] = type_
	type_ = &Type{
		"QPaintEvent*",
		classes["QPaintEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[581] = type_
	type_ = &Type{
		"QPainter*",
		classes["QPainter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[582] = type_
	type_ = &Type{
		"QPainter::CompositionMode",
		classes["QPainter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[583] = type_
	type_ = &Type{
		"QPainter::PixmapFragment",
		classes["QPainter::PixmapFragment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[584] = type_
	type_ = &Type{
		"QPainter::PixmapFragment*",
		classes["QPainter::PixmapFragment"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[585] = type_
	type_ = &Type{
		"QPainter::PixmapFragmentHint",
		classes["QPainter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[586] = type_
	type_ = &Type{
		"QPainter::RenderHint",
		classes["QPainter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[587] = type_
	type_ = &Type{
		"QPainterPath",
		classes["QPainterPath"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[588] = type_
	type_ = &Type{
		"QPainterPath&",
		classes["QPainterPath"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[589] = type_
	type_ = &Type{
		"QPainterPath*",
		classes["QPainterPath"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[590] = type_
	type_ = &Type{
		"QPainterPath::Element*",
		classes["QPainterPath::Element"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[591] = type_
	type_ = &Type{
		"QPainterPath::ElementType",
		classes["QPainterPath"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[592] = type_
	type_ = &Type{
		"QPainterPathStroker*",
		classes["QPainterPathStroker"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[593] = type_
	type_ = &Type{
		"QPalette",
		classes["QPalette"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[594] = type_
	type_ = &Type{
		"QPalette&",
		classes["QPalette"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[595] = type_
	type_ = &Type{
		"QPalette*",
		classes["QPalette"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[596] = type_
	type_ = &Type{
		"QPalette::ColorGroup",
		classes["QPalette"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[597] = type_
	type_ = &Type{
		"QPalette::ColorRole",
		classes["QPalette"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[598] = type_
	type_ = &Type{
		"QPanGesture*",
		classes["QPanGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[599] = type_
	type_ = &Type{
		"QPen",
		classes["QPen"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[600] = type_
	type_ = &Type{
		"QPen&",
		classes["QPen"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[601] = type_
	type_ = &Type{
		"QPen*",
		classes["QPen"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[602] = type_
	type_ = &Type{
		"QPenPrivate*&",
		classes["QPenPrivate"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[603] = type_
	type_ = &Type{
		"QPicture",
		classes["QPicture"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[604] = type_
	type_ = &Type{
		"QPicture&",
		classes["QPicture"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[605] = type_
	type_ = &Type{
		"QPicture*",
		classes["QPicture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[606] = type_
	type_ = &Type{
		"QPictureFormatInterface*",
		classes["QPictureFormatInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[607] = type_
	type_ = &Type{
		"QPictureFormatPlugin*",
		classes["QPictureFormatPlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[608] = type_
	type_ = &Type{
		"QPictureIO*",
		classes["QPictureIO"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[609] = type_
	type_ = &Type{
		"QPinchGesture*",
		classes["QPinchGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[610] = type_
	type_ = &Type{
		"QPinchGesture::ChangeFlag",
		classes["QPinchGesture"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[611] = type_
	type_ = &Type{
		"QPixmap",
		classes["QPixmap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[612] = type_
	type_ = &Type{
		"QPixmap&",
		classes["QPixmap"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[613] = type_
	type_ = &Type{
		"QPixmap*",
		classes["QPixmap"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[614] = type_
	type_ = &Type{
		"QPixmap::ShareMode",
		classes["QPixmap"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[615] = type_
	type_ = &Type{
		"QPixmap::Type",
		classes["QPixmap"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[616] = type_
	type_ = &Type{
		"QPixmapCache*",
		classes["QPixmapCache"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[617] = type_
	type_ = &Type{
		"QPixmapCache::Key",
		classes["QPixmapCache::Key"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[618] = type_
	type_ = &Type{
		"QPixmapCache::Key&",
		classes["QPixmapCache::Key"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[619] = type_
	type_ = &Type{
		"QPixmapCache::Key*",
		classes["QPixmapCache::Key"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[620] = type_
	type_ = &Type{
		"QPlainTextDocumentLayout*",
		classes["QPlainTextDocumentLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[621] = type_
	type_ = &Type{
		"QPlainTextEdit*",
		classes["QPlainTextEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[622] = type_
	type_ = &Type{
		"QPlainTextEdit::LineWrapMode",
		classes["QPlainTextEdit"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[623] = type_
	type_ = &Type{
		"QPoint",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[624] = type_
	type_ = &Type{
		"QPoint&",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[625] = type_
	type_ = &Type{
		"QPoint*",
		classes["QPoint"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[626] = type_
	type_ = &Type{
		"QPointF",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[627] = type_
	type_ = &Type{
		"QPointF&",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[628] = type_
	type_ = &Type{
		"QPointF*",
		classes["QPointF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[629] = type_
	type_ = &Type{
		"QPolygon",
		classes["QPolygon"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[630] = type_
	type_ = &Type{
		"QPolygon&",
		classes["QPolygon"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[631] = type_
	type_ = &Type{
		"QPolygon*",
		classes["QPolygon"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[632] = type_
	type_ = &Type{
		"QPolygonF",
		classes["QPolygonF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[633] = type_
	type_ = &Type{
		"QPolygonF&",
		classes["QPolygonF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[634] = type_
	type_ = &Type{
		"QPolygonF*",
		classes["QPolygonF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[635] = type_
	type_ = &Type{
		"QPostEventList*",
		classes["QPostEventList"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[636] = type_
	type_ = &Type{
		"QPrintDialog*",
		classes["QPrintDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[637] = type_
	type_ = &Type{
		"QPrintEngine*",
		classes["QPrintEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[638] = type_
	type_ = &Type{
		"QPrintEngine::PrintEnginePropertyKey",
		classes["QPrintEngine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[639] = type_
	type_ = &Type{
		"QPrintPreviewDialog*",
		classes["QPrintPreviewDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[640] = type_
	type_ = &Type{
		"QPrintPreviewWidget*",
		classes["QPrintPreviewWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[641] = type_
	type_ = &Type{
		"QPrintPreviewWidget::ViewMode",
		classes["QPrintPreviewWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[642] = type_
	type_ = &Type{
		"QPrintPreviewWidget::ZoomMode",
		classes["QPrintPreviewWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[643] = type_
	type_ = &Type{
		"QPrinter*",
		classes["QPrinter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[644] = type_
	type_ = &Type{
		"QPrinter::ColorMode",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[645] = type_
	type_ = &Type{
		"QPrinter::DuplexMode",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[646] = type_
	type_ = &Type{
		"QPrinter::Orientation",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[647] = type_
	type_ = &Type{
		"QPrinter::OutputFormat",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[648] = type_
	type_ = &Type{
		"QPrinter::PageOrder",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[649] = type_
	type_ = &Type{
		"QPrinter::PageSize",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[650] = type_
	type_ = &Type{
		"QPrinter::PaperSource",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[651] = type_
	type_ = &Type{
		"QPrinter::PrintRange",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[652] = type_
	type_ = &Type{
		"QPrinter::PrinterMode",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[653] = type_
	type_ = &Type{
		"QPrinter::PrinterState",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[654] = type_
	type_ = &Type{
		"QPrinter::Unit",
		classes["QPrinter"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[655] = type_
	type_ = &Type{
		"QPrinterInfo",
		classes["QPrinterInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[656] = type_
	type_ = &Type{
		"QPrinterInfo&",
		classes["QPrinterInfo"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[657] = type_
	type_ = &Type{
		"QPrinterInfo*",
		classes["QPrinterInfo"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[658] = type_
	type_ = &Type{
		"QProgressBar*",
		classes["QProgressBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[659] = type_
	type_ = &Type{
		"QProgressBar::Direction",
		classes["QProgressBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[660] = type_
	type_ = &Type{
		"QProgressDialog*",
		classes["QProgressDialog"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[661] = type_
	type_ = &Type{
		"QProxyModel*",
		classes["QProxyModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[662] = type_
	type_ = &Type{
		"QProxyStyle*",
		classes["QProxyStyle"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[663] = type_
	type_ = &Type{
		"QPushButton*",
		classes["QPushButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[664] = type_
	type_ = &Type{
		"QQuaternion",
		classes["QQuaternion"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[665] = type_
	type_ = &Type{
		"QQuaternion&",
		classes["QQuaternion"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[666] = type_
	type_ = &Type{
		"QQuaternion*",
		classes["QQuaternion"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[667] = type_
	type_ = &Type{
		"QRadialGradient*",
		classes["QRadialGradient"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[668] = type_
	type_ = &Type{
		"QRadioButton*",
		classes["QRadioButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[669] = type_
	type_ = &Type{
		"QRawFont",
		classes["QRawFont"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[670] = type_
	type_ = &Type{
		"QRawFont&",
		classes["QRawFont"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[671] = type_
	type_ = &Type{
		"QRawFont*",
		classes["QRawFont"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[672] = type_
	type_ = &Type{
		"QRawFont::AntialiasingType",
		classes["QRawFont"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[673] = type_
	type_ = &Type{
		"QRect",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[674] = type_
	type_ = &Type{
		"QRect&",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[675] = type_
	type_ = &Type{
		"QRect*",
		classes["QRect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[676] = type_
	type_ = &Type{
		"QRectF",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[677] = type_
	type_ = &Type{
		"QRectF&",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[678] = type_
	type_ = &Type{
		"QRectF*",
		classes["QRectF"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[679] = type_
	type_ = &Type{
		"QRegExp",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[680] = type_
	type_ = &Type{
		"QRegExp&",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[681] = type_
	type_ = &Type{
		"QRegExpValidator*",
		classes["QRegExpValidator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[682] = type_
	type_ = &Type{
		"QRegion",
		classes["QRegion"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[683] = type_
	type_ = &Type{
		"QRegion&",
		classes["QRegion"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[684] = type_
	type_ = &Type{
		"QRegion*",
		classes["QRegion"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[685] = type_
	type_ = &Type{
		"QRegion::RegionType",
		classes["QRegion"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[686] = type_
	type_ = &Type{
		"QResizeEvent*",
		classes["QResizeEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[687] = type_
	type_ = &Type{
		"QRubberBand*",
		classes["QRubberBand"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[688] = type_
	type_ = &Type{
		"QRubberBand::Shape",
		classes["QRubberBand"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[689] = type_
	type_ = &Type{
		"QScopedPointer<QBrushData,QBrushDataPointerDeleter>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[690] = type_
	type_ = &Type{
		"QScrollArea*",
		classes["QScrollArea"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[691] = type_
	type_ = &Type{
		"QScrollBar*",
		classes["QScrollBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[692] = type_
	type_ = &Type{
		"QSessionManager&",
		classes["QSessionManager"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[693] = type_
	type_ = &Type{
		"QSessionManager::RestartHint",
		classes["QSessionManager"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[694] = type_
	type_ = &Type{
		"QSet<QAccessible::Method>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[695] = type_
	type_ = &Type{
		"QShortcut*",
		classes["QShortcut"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[696] = type_
	type_ = &Type{
		"QShortcutEvent*",
		classes["QShortcutEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[697] = type_
	type_ = &Type{
		"QShowEvent*",
		classes["QShowEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[698] = type_
	type_ = &Type{
		"QSize",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[699] = type_
	type_ = &Type{
		"QSize&",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[700] = type_
	type_ = &Type{
		"QSizeF",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[701] = type_
	type_ = &Type{
		"QSizeF&",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[702] = type_
	type_ = &Type{
		"QSizeGrip*",
		classes["QSizeGrip"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[703] = type_
	type_ = &Type{
		"QSizePolicy",
		classes["QSizePolicy"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[704] = type_
	type_ = &Type{
		"QSizePolicy&",
		classes["QSizePolicy"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[705] = type_
	type_ = &Type{
		"QSizePolicy*",
		classes["QSizePolicy"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[706] = type_
	type_ = &Type{
		"QSizePolicy::ControlType",
		classes["QSizePolicy"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[707] = type_
	type_ = &Type{
		"QSizePolicy::Policy",
		classes["QSizePolicy"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[708] = type_
	type_ = &Type{
		"QSizePolicy::PolicyFlag",
		classes["QSizePolicy"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[709] = type_
	type_ = &Type{
		"QSizePolicy::SizePolicyMasks",
		classes["QSizePolicy"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[710] = type_
	type_ = &Type{
		"QSlider*",
		classes["QSlider"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[711] = type_
	type_ = &Type{
		"QSlider::TickPosition",
		classes["QSlider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[712] = type_
	type_ = &Type{
		"QSortFilterProxyModel*",
		classes["QSortFilterProxyModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[713] = type_
	type_ = &Type{
		"QSound*",
		classes["QSound"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[714] = type_
	type_ = &Type{
		"QSpacerItem*",
		classes["QSpacerItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[715] = type_
	type_ = &Type{
		"QSpinBox*",
		classes["QSpinBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[716] = type_
	type_ = &Type{
		"QSplashScreen*",
		classes["QSplashScreen"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[717] = type_
	type_ = &Type{
		"QSplitter&",
		classes["QSplitter"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[718] = type_
	type_ = &Type{
		"QSplitter*",
		classes["QSplitter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[719] = type_
	type_ = &Type{
		"QSplitterHandle*",
		classes["QSplitterHandle"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[720] = type_
	type_ = &Type{
		"QStackedLayout*",
		classes["QStackedLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[721] = type_
	type_ = &Type{
		"QStackedLayout::StackingMode",
		classes["QStackedLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[722] = type_
	type_ = &Type{
		"QStackedWidget*",
		classes["QStackedWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[723] = type_
	type_ = &Type{
		"QStandardItem&",
		classes["QStandardItem"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[724] = type_
	type_ = &Type{
		"QStandardItem*",
		classes["QStandardItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[725] = type_
	type_ = &Type{
		"QStandardItem::ItemType",
		classes["QStandardItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[726] = type_
	type_ = &Type{
		"QStandardItemModel*",
		classes["QStandardItemModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[727] = type_
	type_ = &Type{
		"QState*",
		classes["QState"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[728] = type_
	type_ = &Type{
		"QStaticText&",
		classes["QStaticText"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[729] = type_
	type_ = &Type{
		"QStaticText*",
		classes["QStaticText"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[730] = type_
	type_ = &Type{
		"QStaticText::PerformanceHint",
		classes["QStaticText"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[731] = type_
	type_ = &Type{
		"QStatusBar*",
		classes["QStatusBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[732] = type_
	type_ = &Type{
		"QStatusTipEvent*",
		classes["QStatusTipEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[733] = type_
	type_ = &Type{
		"QString",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[734] = type_
	type_ = &Type{
		"QString&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[735] = type_
	type_ = &Type{
		"QString*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[736] = type_
	type_ = &Type{
		"QString::Null",
		classes["QString::Null"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[737] = type_
	type_ = &Type{
		"QString::SectionFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[738] = type_
	type_ = &Type{
		"QStringList",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[739] = type_
	type_ = &Type{
		"QStringList&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[740] = type_
	type_ = &Type{
		"QStringList*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[741] = type_
	type_ = &Type{
		"QStringListModel*",
		classes["QStringListModel"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[742] = type_
	type_ = &Type{
		"QStyle&",
		classes["QStyle"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[743] = type_
	type_ = &Type{
		"QStyle*",
		classes["QStyle"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[744] = type_
	type_ = &Type{
		"QStyle::ComplexControl",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[745] = type_
	type_ = &Type{
		"QStyle::ContentsType",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[746] = type_
	type_ = &Type{
		"QStyle::ControlElement",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[747] = type_
	type_ = &Type{
		"QStyle::PixelMetric",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[748] = type_
	type_ = &Type{
		"QStyle::PrimitiveElement",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[749] = type_
	type_ = &Type{
		"QStyle::RequestSoftwareInputPanel",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[750] = type_
	type_ = &Type{
		"QStyle::StandardPixmap",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[751] = type_
	type_ = &Type{
		"QStyle::StateFlag",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[752] = type_
	type_ = &Type{
		"QStyle::StyleHint",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[753] = type_
	type_ = &Type{
		"QStyle::SubControl",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[754] = type_
	type_ = &Type{
		"QStyle::SubElement",
		classes["QStyle"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[755] = type_
	type_ = &Type{
		"QStyleFactory*",
		classes["QStyleFactory"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[756] = type_
	type_ = &Type{
		"QStyleFactoryInterface*",
		classes["QStyleFactoryInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[757] = type_
	type_ = &Type{
		"QStyleHintReturn*",
		classes["QStyleHintReturn"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[758] = type_
	type_ = &Type{
		"QStyleHintReturn::HintReturnType",
		classes["QStyleHintReturn"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[759] = type_
	type_ = &Type{
		"QStyleHintReturn::StyleOptionType",
		classes["QStyleHintReturn"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[760] = type_
	type_ = &Type{
		"QStyleHintReturn::StyleOptionVersion",
		classes["QStyleHintReturn"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[761] = type_
	type_ = &Type{
		"QStyleHintReturnMask*",
		classes["QStyleHintReturnMask"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[762] = type_
	type_ = &Type{
		"QStyleHintReturnMask::StyleOptionType",
		classes["QStyleHintReturnMask"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[763] = type_
	type_ = &Type{
		"QStyleHintReturnMask::StyleOptionVersion",
		classes["QStyleHintReturnMask"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[764] = type_
	type_ = &Type{
		"QStyleHintReturnVariant*",
		classes["QStyleHintReturnVariant"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[765] = type_
	type_ = &Type{
		"QStyleHintReturnVariant::StyleOptionType",
		classes["QStyleHintReturnVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[766] = type_
	type_ = &Type{
		"QStyleHintReturnVariant::StyleOptionVersion",
		classes["QStyleHintReturnVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[767] = type_
	type_ = &Type{
		"QStyleOption&",
		classes["QStyleOption"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[768] = type_
	type_ = &Type{
		"QStyleOption*",
		classes["QStyleOption"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[769] = type_
	type_ = &Type{
		"QStyleOption::OptionType",
		classes["QStyleOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[770] = type_
	type_ = &Type{
		"QStyleOption::StyleOptionType",
		classes["QStyleOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[771] = type_
	type_ = &Type{
		"QStyleOption::StyleOptionVersion",
		classes["QStyleOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[772] = type_
	type_ = &Type{
		"QStyleOptionButton*",
		classes["QStyleOptionButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[773] = type_
	type_ = &Type{
		"QStyleOptionButton::ButtonFeature",
		classes["QStyleOptionButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[774] = type_
	type_ = &Type{
		"QStyleOptionButton::StyleOptionType",
		classes["QStyleOptionButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[775] = type_
	type_ = &Type{
		"QStyleOptionButton::StyleOptionVersion",
		classes["QStyleOptionButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[776] = type_
	type_ = &Type{
		"QStyleOptionComboBox*",
		classes["QStyleOptionComboBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[777] = type_
	type_ = &Type{
		"QStyleOptionComboBox::StyleOptionType",
		classes["QStyleOptionComboBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[778] = type_
	type_ = &Type{
		"QStyleOptionComboBox::StyleOptionVersion",
		classes["QStyleOptionComboBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[779] = type_
	type_ = &Type{
		"QStyleOptionComplex*",
		classes["QStyleOptionComplex"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[780] = type_
	type_ = &Type{
		"QStyleOptionComplex::StyleOptionType",
		classes["QStyleOptionComplex"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[781] = type_
	type_ = &Type{
		"QStyleOptionComplex::StyleOptionVersion",
		classes["QStyleOptionComplex"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[782] = type_
	type_ = &Type{
		"QStyleOptionDockWidget*",
		classes["QStyleOptionDockWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[783] = type_
	type_ = &Type{
		"QStyleOptionDockWidget::StyleOptionType",
		classes["QStyleOptionDockWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[784] = type_
	type_ = &Type{
		"QStyleOptionDockWidget::StyleOptionVersion",
		classes["QStyleOptionDockWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[785] = type_
	type_ = &Type{
		"QStyleOptionDockWidgetV2&",
		classes["QStyleOptionDockWidgetV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[786] = type_
	type_ = &Type{
		"QStyleOptionDockWidgetV2*",
		classes["QStyleOptionDockWidgetV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[787] = type_
	type_ = &Type{
		"QStyleOptionDockWidgetV2::StyleOptionVersion",
		classes["QStyleOptionDockWidgetV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[788] = type_
	type_ = &Type{
		"QStyleOptionFocusRect*",
		classes["QStyleOptionFocusRect"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[789] = type_
	type_ = &Type{
		"QStyleOptionFocusRect::StyleOptionType",
		classes["QStyleOptionFocusRect"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[790] = type_
	type_ = &Type{
		"QStyleOptionFocusRect::StyleOptionVersion",
		classes["QStyleOptionFocusRect"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[791] = type_
	type_ = &Type{
		"QStyleOptionFrame*",
		classes["QStyleOptionFrame"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[792] = type_
	type_ = &Type{
		"QStyleOptionFrame::StyleOptionType",
		classes["QStyleOptionFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[793] = type_
	type_ = &Type{
		"QStyleOptionFrame::StyleOptionVersion",
		classes["QStyleOptionFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[794] = type_
	type_ = &Type{
		"QStyleOptionFrameV2&",
		classes["QStyleOptionFrameV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[795] = type_
	type_ = &Type{
		"QStyleOptionFrameV2*",
		classes["QStyleOptionFrameV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[796] = type_
	type_ = &Type{
		"QStyleOptionFrameV2::FrameFeature",
		classes["QStyleOptionFrameV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[797] = type_
	type_ = &Type{
		"QStyleOptionFrameV2::StyleOptionVersion",
		classes["QStyleOptionFrameV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[798] = type_
	type_ = &Type{
		"QStyleOptionFrameV3&",
		classes["QStyleOptionFrameV3"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[799] = type_
	type_ = &Type{
		"QStyleOptionFrameV3*",
		classes["QStyleOptionFrameV3"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[800] = type_
	type_ = &Type{
		"QStyleOptionFrameV3::StyleOptionVersion",
		classes["QStyleOptionFrameV3"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[801] = type_
	type_ = &Type{
		"QStyleOptionGraphicsItem*",
		classes["QStyleOptionGraphicsItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[802] = type_
	type_ = &Type{
		"QStyleOptionGraphicsItem::StyleOptionType",
		classes["QStyleOptionGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[803] = type_
	type_ = &Type{
		"QStyleOptionGraphicsItem::StyleOptionVersion",
		classes["QStyleOptionGraphicsItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[804] = type_
	type_ = &Type{
		"QStyleOptionGroupBox*",
		classes["QStyleOptionGroupBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[805] = type_
	type_ = &Type{
		"QStyleOptionGroupBox::StyleOptionType",
		classes["QStyleOptionGroupBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[806] = type_
	type_ = &Type{
		"QStyleOptionGroupBox::StyleOptionVersion",
		classes["QStyleOptionGroupBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[807] = type_
	type_ = &Type{
		"QStyleOptionHeader*",
		classes["QStyleOptionHeader"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[808] = type_
	type_ = &Type{
		"QStyleOptionHeader::SectionPosition",
		classes["QStyleOptionHeader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[809] = type_
	type_ = &Type{
		"QStyleOptionHeader::SelectedPosition",
		classes["QStyleOptionHeader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[810] = type_
	type_ = &Type{
		"QStyleOptionHeader::SortIndicator",
		classes["QStyleOptionHeader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[811] = type_
	type_ = &Type{
		"QStyleOptionHeader::StyleOptionType",
		classes["QStyleOptionHeader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[812] = type_
	type_ = &Type{
		"QStyleOptionHeader::StyleOptionVersion",
		classes["QStyleOptionHeader"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[813] = type_
	type_ = &Type{
		"QStyleOptionMenuItem*",
		classes["QStyleOptionMenuItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[814] = type_
	type_ = &Type{
		"QStyleOptionMenuItem::CheckType",
		classes["QStyleOptionMenuItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[815] = type_
	type_ = &Type{
		"QStyleOptionMenuItem::MenuItemType",
		classes["QStyleOptionMenuItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[816] = type_
	type_ = &Type{
		"QStyleOptionMenuItem::StyleOptionType",
		classes["QStyleOptionMenuItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[817] = type_
	type_ = &Type{
		"QStyleOptionMenuItem::StyleOptionVersion",
		classes["QStyleOptionMenuItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[818] = type_
	type_ = &Type{
		"QStyleOptionProgressBar*",
		classes["QStyleOptionProgressBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[819] = type_
	type_ = &Type{
		"QStyleOptionProgressBar::StyleOptionType",
		classes["QStyleOptionProgressBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[820] = type_
	type_ = &Type{
		"QStyleOptionProgressBar::StyleOptionVersion",
		classes["QStyleOptionProgressBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[821] = type_
	type_ = &Type{
		"QStyleOptionProgressBarV2&",
		classes["QStyleOptionProgressBarV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[822] = type_
	type_ = &Type{
		"QStyleOptionProgressBarV2*",
		classes["QStyleOptionProgressBarV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[823] = type_
	type_ = &Type{
		"QStyleOptionProgressBarV2::StyleOptionType",
		classes["QStyleOptionProgressBarV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[824] = type_
	type_ = &Type{
		"QStyleOptionProgressBarV2::StyleOptionVersion",
		classes["QStyleOptionProgressBarV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[825] = type_
	type_ = &Type{
		"QStyleOptionQ3ListViewItem::Q3ListViewItemFeature",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[826] = type_
	type_ = &Type{
		"QStyleOptionRubberBand*",
		classes["QStyleOptionRubberBand"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[827] = type_
	type_ = &Type{
		"QStyleOptionRubberBand::StyleOptionType",
		classes["QStyleOptionRubberBand"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[828] = type_
	type_ = &Type{
		"QStyleOptionRubberBand::StyleOptionVersion",
		classes["QStyleOptionRubberBand"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[829] = type_
	type_ = &Type{
		"QStyleOptionSizeGrip*",
		classes["QStyleOptionSizeGrip"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[830] = type_
	type_ = &Type{
		"QStyleOptionSizeGrip::StyleOptionType",
		classes["QStyleOptionSizeGrip"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[831] = type_
	type_ = &Type{
		"QStyleOptionSizeGrip::StyleOptionVersion",
		classes["QStyleOptionSizeGrip"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[832] = type_
	type_ = &Type{
		"QStyleOptionSlider*",
		classes["QStyleOptionSlider"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[833] = type_
	type_ = &Type{
		"QStyleOptionSlider::StyleOptionType",
		classes["QStyleOptionSlider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[834] = type_
	type_ = &Type{
		"QStyleOptionSlider::StyleOptionVersion",
		classes["QStyleOptionSlider"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[835] = type_
	type_ = &Type{
		"QStyleOptionSpinBox*",
		classes["QStyleOptionSpinBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[836] = type_
	type_ = &Type{
		"QStyleOptionSpinBox::StyleOptionType",
		classes["QStyleOptionSpinBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[837] = type_
	type_ = &Type{
		"QStyleOptionSpinBox::StyleOptionVersion",
		classes["QStyleOptionSpinBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[838] = type_
	type_ = &Type{
		"QStyleOptionTab*",
		classes["QStyleOptionTab"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[839] = type_
	type_ = &Type{
		"QStyleOptionTab::CornerWidget",
		classes["QStyleOptionTab"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[840] = type_
	type_ = &Type{
		"QStyleOptionTab::SelectedPosition",
		classes["QStyleOptionTab"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[841] = type_
	type_ = &Type{
		"QStyleOptionTab::StyleOptionType",
		classes["QStyleOptionTab"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[842] = type_
	type_ = &Type{
		"QStyleOptionTab::StyleOptionVersion",
		classes["QStyleOptionTab"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[843] = type_
	type_ = &Type{
		"QStyleOptionTab::TabPosition",
		classes["QStyleOptionTab"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[844] = type_
	type_ = &Type{
		"QStyleOptionTabBarBase*",
		classes["QStyleOptionTabBarBase"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[845] = type_
	type_ = &Type{
		"QStyleOptionTabBarBase::StyleOptionType",
		classes["QStyleOptionTabBarBase"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[846] = type_
	type_ = &Type{
		"QStyleOptionTabBarBase::StyleOptionVersion",
		classes["QStyleOptionTabBarBase"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[847] = type_
	type_ = &Type{
		"QStyleOptionTabBarBaseV2&",
		classes["QStyleOptionTabBarBaseV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[848] = type_
	type_ = &Type{
		"QStyleOptionTabBarBaseV2*",
		classes["QStyleOptionTabBarBaseV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[849] = type_
	type_ = &Type{
		"QStyleOptionTabBarBaseV2::StyleOptionVersion",
		classes["QStyleOptionTabBarBaseV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[850] = type_
	type_ = &Type{
		"QStyleOptionTabV2&",
		classes["QStyleOptionTabV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[851] = type_
	type_ = &Type{
		"QStyleOptionTabV2*",
		classes["QStyleOptionTabV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[852] = type_
	type_ = &Type{
		"QStyleOptionTabV2::StyleOptionVersion",
		classes["QStyleOptionTabV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[853] = type_
	type_ = &Type{
		"QStyleOptionTabV3&",
		classes["QStyleOptionTabV3"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[854] = type_
	type_ = &Type{
		"QStyleOptionTabV3*",
		classes["QStyleOptionTabV3"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[855] = type_
	type_ = &Type{
		"QStyleOptionTabV3::StyleOptionVersion",
		classes["QStyleOptionTabV3"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[856] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrame*",
		classes["QStyleOptionTabWidgetFrame"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[857] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrame::StyleOptionType",
		classes["QStyleOptionTabWidgetFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[858] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrame::StyleOptionVersion",
		classes["QStyleOptionTabWidgetFrame"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[859] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrameV2&",
		classes["QStyleOptionTabWidgetFrameV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[860] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrameV2*",
		classes["QStyleOptionTabWidgetFrameV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[861] = type_
	type_ = &Type{
		"QStyleOptionTabWidgetFrameV2::StyleOptionVersion",
		classes["QStyleOptionTabWidgetFrameV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[862] = type_
	type_ = &Type{
		"QStyleOptionTitleBar*",
		classes["QStyleOptionTitleBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[863] = type_
	type_ = &Type{
		"QStyleOptionTitleBar::StyleOptionType",
		classes["QStyleOptionTitleBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[864] = type_
	type_ = &Type{
		"QStyleOptionTitleBar::StyleOptionVersion",
		classes["QStyleOptionTitleBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[865] = type_
	type_ = &Type{
		"QStyleOptionToolBar*",
		classes["QStyleOptionToolBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[866] = type_
	type_ = &Type{
		"QStyleOptionToolBar::StyleOptionType",
		classes["QStyleOptionToolBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[867] = type_
	type_ = &Type{
		"QStyleOptionToolBar::StyleOptionVersion",
		classes["QStyleOptionToolBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[868] = type_
	type_ = &Type{
		"QStyleOptionToolBar::ToolBarFeature",
		classes["QStyleOptionToolBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[869] = type_
	type_ = &Type{
		"QStyleOptionToolBar::ToolBarPosition",
		classes["QStyleOptionToolBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[870] = type_
	type_ = &Type{
		"QStyleOptionToolBox*",
		classes["QStyleOptionToolBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[871] = type_
	type_ = &Type{
		"QStyleOptionToolBox::StyleOptionType",
		classes["QStyleOptionToolBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[872] = type_
	type_ = &Type{
		"QStyleOptionToolBox::StyleOptionVersion",
		classes["QStyleOptionToolBox"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[873] = type_
	type_ = &Type{
		"QStyleOptionToolBoxV2&",
		classes["QStyleOptionToolBoxV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[874] = type_
	type_ = &Type{
		"QStyleOptionToolBoxV2*",
		classes["QStyleOptionToolBoxV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[875] = type_
	type_ = &Type{
		"QStyleOptionToolBoxV2::SelectedPosition",
		classes["QStyleOptionToolBoxV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[876] = type_
	type_ = &Type{
		"QStyleOptionToolBoxV2::StyleOptionVersion",
		classes["QStyleOptionToolBoxV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[877] = type_
	type_ = &Type{
		"QStyleOptionToolBoxV2::TabPosition",
		classes["QStyleOptionToolBoxV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[878] = type_
	type_ = &Type{
		"QStyleOptionToolButton*",
		classes["QStyleOptionToolButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[879] = type_
	type_ = &Type{
		"QStyleOptionToolButton::StyleOptionType",
		classes["QStyleOptionToolButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[880] = type_
	type_ = &Type{
		"QStyleOptionToolButton::StyleOptionVersion",
		classes["QStyleOptionToolButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[881] = type_
	type_ = &Type{
		"QStyleOptionToolButton::ToolButtonFeature",
		classes["QStyleOptionToolButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[882] = type_
	type_ = &Type{
		"QStyleOptionViewItem",
		classes["QStyleOptionViewItem"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[883] = type_
	type_ = &Type{
		"QStyleOptionViewItem*",
		classes["QStyleOptionViewItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[884] = type_
	type_ = &Type{
		"QStyleOptionViewItem::Position",
		classes["QStyleOptionViewItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[885] = type_
	type_ = &Type{
		"QStyleOptionViewItem::StyleOptionType",
		classes["QStyleOptionViewItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[886] = type_
	type_ = &Type{
		"QStyleOptionViewItem::StyleOptionVersion",
		classes["QStyleOptionViewItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[887] = type_
	type_ = &Type{
		"QStyleOptionViewItemV2&",
		classes["QStyleOptionViewItemV2"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[888] = type_
	type_ = &Type{
		"QStyleOptionViewItemV2*",
		classes["QStyleOptionViewItemV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[889] = type_
	type_ = &Type{
		"QStyleOptionViewItemV2::StyleOptionVersion",
		classes["QStyleOptionViewItemV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[890] = type_
	type_ = &Type{
		"QStyleOptionViewItemV2::ViewItemFeature",
		classes["QStyleOptionViewItemV2"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[891] = type_
	type_ = &Type{
		"QStyleOptionViewItemV3&",
		classes["QStyleOptionViewItemV3"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[892] = type_
	type_ = &Type{
		"QStyleOptionViewItemV3*",
		classes["QStyleOptionViewItemV3"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[893] = type_
	type_ = &Type{
		"QStyleOptionViewItemV3::StyleOptionVersion",
		classes["QStyleOptionViewItemV3"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[894] = type_
	type_ = &Type{
		"QStyleOptionViewItemV4&",
		classes["QStyleOptionViewItemV4"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[895] = type_
	type_ = &Type{
		"QStyleOptionViewItemV4*",
		classes["QStyleOptionViewItemV4"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[896] = type_
	type_ = &Type{
		"QStyleOptionViewItemV4::StyleOptionVersion",
		classes["QStyleOptionViewItemV4"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[897] = type_
	type_ = &Type{
		"QStyleOptionViewItemV4::ViewItemPosition",
		classes["QStyleOptionViewItemV4"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[898] = type_
	type_ = &Type{
		"QStylePainter*",
		classes["QStylePainter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[899] = type_
	type_ = &Type{
		"QStylePlugin*",
		classes["QStylePlugin"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[900] = type_
	type_ = &Type{
		"QStyledItemDelegate*",
		classes["QStyledItemDelegate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[901] = type_
	type_ = &Type{
		"QSwipeGesture*",
		classes["QSwipeGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[902] = type_
	type_ = &Type{
		"QSwipeGesture::SwipeDirection",
		classes["QSwipeGesture"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[903] = type_
	type_ = &Type{
		"QSyntaxHighlighter*",
		classes["QSyntaxHighlighter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[904] = type_
	type_ = &Type{
		"QSystemTrayIcon*",
		classes["QSystemTrayIcon"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[905] = type_
	type_ = &Type{
		"QSystemTrayIcon::ActivationReason",
		classes["QSystemTrayIcon"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[906] = type_
	type_ = &Type{
		"QSystemTrayIcon::MessageIcon",
		classes["QSystemTrayIcon"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[907] = type_
	type_ = &Type{
		"QTabBar*",
		classes["QTabBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[908] = type_
	type_ = &Type{
		"QTabBar::ButtonPosition",
		classes["QTabBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[909] = type_
	type_ = &Type{
		"QTabBar::SelectionBehavior",
		classes["QTabBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[910] = type_
	type_ = &Type{
		"QTabBar::Shape",
		classes["QTabBar"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[911] = type_
	type_ = &Type{
		"QTabWidget*",
		classes["QTabWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[912] = type_
	type_ = &Type{
		"QTabWidget::TabPosition",
		classes["QTabWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[913] = type_
	type_ = &Type{
		"QTabWidget::TabShape",
		classes["QTabWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[914] = type_
	type_ = &Type{
		"QTableView*",
		classes["QTableView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[915] = type_
	type_ = &Type{
		"QTableWidget*",
		classes["QTableWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[916] = type_
	type_ = &Type{
		"QTableWidgetItem&",
		classes["QTableWidgetItem"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[917] = type_
	type_ = &Type{
		"QTableWidgetItem*",
		classes["QTableWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[918] = type_
	type_ = &Type{
		"QTableWidgetItem::ItemType",
		classes["QTableWidgetItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[919] = type_
	type_ = &Type{
		"QTableWidgetSelectionRange*",
		classes["QTableWidgetSelectionRange"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[920] = type_
	type_ = &Type{
		"QTabletEvent*",
		classes["QTabletEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[921] = type_
	type_ = &Type{
		"QTabletEvent::PointerType",
		classes["QTabletEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[922] = type_
	type_ = &Type{
		"QTabletEvent::TabletDevice",
		classes["QTabletEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[923] = type_
	type_ = &Type{
		"QTapAndHoldGesture*",
		classes["QTapAndHoldGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[924] = type_
	type_ = &Type{
		"QTapGesture*",
		classes["QTapGesture"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[925] = type_
	type_ = &Type{
		"QTextBlock",
		classes["QTextBlock"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[926] = type_
	type_ = &Type{
		"QTextBlock&",
		classes["QTextBlock"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[927] = type_
	type_ = &Type{
		"QTextBlock*",
		classes["QTextBlock"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[928] = type_
	type_ = &Type{
		"QTextBlock::iterator",
		classes["QTextBlock::iterator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[929] = type_
	type_ = &Type{
		"QTextBlock::iterator&",
		classes["QTextBlock::iterator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[930] = type_
	type_ = &Type{
		"QTextBlock::iterator*",
		classes["QTextBlock::iterator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[931] = type_
	type_ = &Type{
		"QTextBlockFormat",
		classes["QTextBlockFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[932] = type_
	type_ = &Type{
		"QTextBlockFormat*",
		classes["QTextBlockFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[933] = type_
	type_ = &Type{
		"QTextBlockFormat::LineHeightTypes",
		classes["QTextBlockFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[934] = type_
	type_ = &Type{
		"QTextBlockGroup*",
		classes["QTextBlockGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[935] = type_
	type_ = &Type{
		"QTextBlockUserData*",
		classes["QTextBlockUserData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[936] = type_
	type_ = &Type{
		"QTextBrowser*",
		classes["QTextBrowser"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[937] = type_
	type_ = &Type{
		"QTextCharFormat",
		classes["QTextCharFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[938] = type_
	type_ = &Type{
		"QTextCharFormat&",
		classes["QTextCharFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[939] = type_
	type_ = &Type{
		"QTextCharFormat*",
		classes["QTextCharFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[940] = type_
	type_ = &Type{
		"QTextCharFormat::UnderlineStyle",
		classes["QTextCharFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[941] = type_
	type_ = &Type{
		"QTextCharFormat::VerticalAlignment",
		classes["QTextCharFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[942] = type_
	type_ = &Type{
		"QTextCodec*",
		classes["QTextCodec"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[943] = type_
	type_ = &Type{
		"QTextCodec::ConversionFlag",
		classes["QTextCodec"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[944] = type_
	type_ = &Type{
		"QTextCursor",
		classes["QTextCursor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[945] = type_
	type_ = &Type{
		"QTextCursor&",
		classes["QTextCursor"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[946] = type_
	type_ = &Type{
		"QTextCursor*",
		classes["QTextCursor"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[947] = type_
	type_ = &Type{
		"QTextCursor::MoveMode",
		classes["QTextCursor"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[948] = type_
	type_ = &Type{
		"QTextCursor::MoveOperation",
		classes["QTextCursor"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[949] = type_
	type_ = &Type{
		"QTextCursor::SelectionType",
		classes["QTextCursor"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[950] = type_
	type_ = &Type{
		"QTextDocument*",
		classes["QTextDocument"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[951] = type_
	type_ = &Type{
		"QTextDocument::FindFlag",
		classes["QTextDocument"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[952] = type_
	type_ = &Type{
		"QTextDocument::MetaInformation",
		classes["QTextDocument"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[953] = type_
	type_ = &Type{
		"QTextDocument::ResourceType",
		classes["QTextDocument"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[954] = type_
	type_ = &Type{
		"QTextDocument::Stacks",
		classes["QTextDocument"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[955] = type_
	type_ = &Type{
		"QTextDocumentFragment",
		classes["QTextDocumentFragment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[956] = type_
	type_ = &Type{
		"QTextDocumentFragment&",
		classes["QTextDocumentFragment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[957] = type_
	type_ = &Type{
		"QTextDocumentFragment*",
		classes["QTextDocumentFragment"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[958] = type_
	type_ = &Type{
		"QTextDocumentPrivate*",
		classes["QTextDocumentPrivate"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[959] = type_
	type_ = &Type{
		"QTextDocumentWriter*",
		classes["QTextDocumentWriter"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[960] = type_
	type_ = &Type{
		"QTextEdit*",
		classes["QTextEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[961] = type_
	type_ = &Type{
		"QTextEdit::AutoFormattingFlag",
		classes["QTextEdit"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[962] = type_
	type_ = &Type{
		"QTextEdit::ExtraSelection*",
		classes["QTextEdit::ExtraSelection"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[963] = type_
	type_ = &Type{
		"QTextEdit::LineWrapMode",
		classes["QTextEdit"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[964] = type_
	type_ = &Type{
		"QTextEngine*",
		classes["QTextEngine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[965] = type_
	type_ = &Type{
		"QTextFormat",
		classes["QTextFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[966] = type_
	type_ = &Type{
		"QTextFormat&",
		classes["QTextFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[967] = type_
	type_ = &Type{
		"QTextFormat*",
		classes["QTextFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[968] = type_
	type_ = &Type{
		"QTextFormat::FormatType",
		classes["QTextFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[969] = type_
	type_ = &Type{
		"QTextFormat::ObjectTypes",
		classes["QTextFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[970] = type_
	type_ = &Type{
		"QTextFormat::PageBreakFlag",
		classes["QTextFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[971] = type_
	type_ = &Type{
		"QTextFormat::Property",
		classes["QTextFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[972] = type_
	type_ = &Type{
		"QTextFragment",
		classes["QTextFragment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[973] = type_
	type_ = &Type{
		"QTextFragment&",
		classes["QTextFragment"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[974] = type_
	type_ = &Type{
		"QTextFragment*",
		classes["QTextFragment"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[975] = type_
	type_ = &Type{
		"QTextFrame*",
		classes["QTextFrame"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[976] = type_
	type_ = &Type{
		"QTextFrame::iterator",
		classes["QTextFrame::iterator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[977] = type_
	type_ = &Type{
		"QTextFrame::iterator&",
		classes["QTextFrame::iterator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[978] = type_
	type_ = &Type{
		"QTextFrame::iterator*",
		classes["QTextFrame::iterator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[979] = type_
	type_ = &Type{
		"QTextFrameFormat",
		classes["QTextFrameFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[980] = type_
	type_ = &Type{
		"QTextFrameFormat*",
		classes["QTextFrameFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[981] = type_
	type_ = &Type{
		"QTextFrameFormat::BorderStyle",
		classes["QTextFrameFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[982] = type_
	type_ = &Type{
		"QTextFrameFormat::Position",
		classes["QTextFrameFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[983] = type_
	type_ = &Type{
		"QTextFrameLayoutData*",
		classes["QTextFrameLayoutData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[984] = type_
	type_ = &Type{
		"QTextImageFormat",
		classes["QTextImageFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[985] = type_
	type_ = &Type{
		"QTextImageFormat*",
		classes["QTextImageFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[986] = type_
	type_ = &Type{
		"QTextInlineObject",
		classes["QTextInlineObject"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[987] = type_
	type_ = &Type{
		"QTextInlineObject*",
		classes["QTextInlineObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[988] = type_
	type_ = &Type{
		"QTextItem*",
		classes["QTextItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[989] = type_
	type_ = &Type{
		"QTextItem::RenderFlag",
		classes["QTextItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[990] = type_
	type_ = &Type{
		"QTextLayout*",
		classes["QTextLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[991] = type_
	type_ = &Type{
		"QTextLayout::CursorMode",
		classes["QTextLayout"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[992] = type_
	type_ = &Type{
		"QTextLayout::FormatRange*",
		classes["QTextLayout::FormatRange"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[993] = type_
	type_ = &Type{
		"QTextLength",
		classes["QTextLength"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[994] = type_
	type_ = &Type{
		"QTextLength&",
		classes["QTextLength"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[995] = type_
	type_ = &Type{
		"QTextLength*",
		classes["QTextLength"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[996] = type_
	type_ = &Type{
		"QTextLength::Type",
		classes["QTextLength"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[997] = type_
	type_ = &Type{
		"QTextLine",
		classes["QTextLine"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[998] = type_
	type_ = &Type{
		"QTextLine*",
		classes["QTextLine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[999] = type_
	type_ = &Type{
		"QTextLine::CursorPosition",
		classes["QTextLine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1000] = type_
	type_ = &Type{
		"QTextLine::Edge",
		classes["QTextLine"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1001] = type_
	type_ = &Type{
		"QTextList*",
		classes["QTextList"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1002] = type_
	type_ = &Type{
		"QTextListFormat",
		classes["QTextListFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1003] = type_
	type_ = &Type{
		"QTextListFormat*",
		classes["QTextListFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1004] = type_
	type_ = &Type{
		"QTextListFormat::Style",
		classes["QTextListFormat"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1005] = type_
	type_ = &Type{
		"QTextObject*",
		classes["QTextObject"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1006] = type_
	type_ = &Type{
		"QTextObjectInterface*",
		classes["QTextObjectInterface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1007] = type_
	type_ = &Type{
		"QTextOption",
		classes["QTextOption"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1008] = type_
	type_ = &Type{
		"QTextOption&",
		classes["QTextOption"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1009] = type_
	type_ = &Type{
		"QTextOption*",
		classes["QTextOption"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1010] = type_
	type_ = &Type{
		"QTextOption::Flag",
		classes["QTextOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1011] = type_
	type_ = &Type{
		"QTextOption::Tab*",
		classes["QTextOption::Tab"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1012] = type_
	type_ = &Type{
		"QTextOption::TabType",
		classes["QTextOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1013] = type_
	type_ = &Type{
		"QTextOption::WrapMode",
		classes["QTextOption"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1014] = type_
	type_ = &Type{
		"QTextStream&",
		classes["QTextStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1015] = type_
	type_ = &Type{
		"QTextStream&(*)(QTextStream&)",
		classes["QTextStream"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1016] = type_
	type_ = &Type{
		"QTextStream::NumberFlag",
		classes["QTextStream"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1017] = type_
	type_ = &Type{
		"QTextStreamManipulator",
		classes["QTextStreamManipulator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1018] = type_
	type_ = &Type{
		"QTextTable*",
		classes["QTextTable"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1019] = type_
	type_ = &Type{
		"QTextTableCell",
		classes["QTextTableCell"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1020] = type_
	type_ = &Type{
		"QTextTableCell&",
		classes["QTextTableCell"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1021] = type_
	type_ = &Type{
		"QTextTableCell*",
		classes["QTextTableCell"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1022] = type_
	type_ = &Type{
		"QTextTableCellFormat",
		classes["QTextTableCellFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1023] = type_
	type_ = &Type{
		"QTextTableCellFormat*",
		classes["QTextTableCellFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1024] = type_
	type_ = &Type{
		"QTextTableFormat",
		classes["QTextTableFormat"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1025] = type_
	type_ = &Type{
		"QTextTableFormat*",
		classes["QTextTableFormat"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1026] = type_
	type_ = &Type{
		"QTileRules*",
		classes["QTileRules"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1027] = type_
	type_ = &Type{
		"QTime",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1028] = type_
	type_ = &Type{
		"QTime&",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1029] = type_
	type_ = &Type{
		"QTimeEdit*",
		classes["QTimeEdit"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1030] = type_
	type_ = &Type{
		"QTimeLine*",
		classes["QTimeLine"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1031] = type_
	type_ = &Type{
		"QTimerEvent*",
		classes["QTimerEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1032] = type_
	type_ = &Type{
		"QToolBar*",
		classes["QToolBar"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1033] = type_
	type_ = &Type{
		"QToolBarChangeEvent*",
		classes["QToolBarChangeEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1034] = type_
	type_ = &Type{
		"QToolBox*",
		classes["QToolBox"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1035] = type_
	type_ = &Type{
		"QToolButton*",
		classes["QToolButton"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1036] = type_
	type_ = &Type{
		"QToolButton::ToolButtonPopupMode",
		classes["QToolButton"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1037] = type_
	type_ = &Type{
		"QToolTip*",
		classes["QToolTip"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1038] = type_
	type_ = &Type{
		"QTouchEvent*",
		classes["QTouchEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1039] = type_
	type_ = &Type{
		"QTouchEvent::DeviceType",
		classes["QTouchEvent"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1040] = type_
	type_ = &Type{
		"QTouchEvent::TouchPoint&",
		classes["QTouchEvent::TouchPoint"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1041] = type_
	type_ = &Type{
		"QTouchEvent::TouchPoint*",
		classes["QTouchEvent::TouchPoint"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1042] = type_
	type_ = &Type{
		"QTransform",
		classes["QTransform"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1043] = type_
	type_ = &Type{
		"QTransform&",
		classes["QTransform"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1044] = type_
	type_ = &Type{
		"QTransform*",
		classes["QTransform"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1045] = type_
	type_ = &Type{
		"QTransform::TransformationType",
		classes["QTransform"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1046] = type_
	type_ = &Type{
		"QTreeView*",
		classes["QTreeView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1047] = type_
	type_ = &Type{
		"QTreeWidget*",
		classes["QTreeWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1048] = type_
	type_ = &Type{
		"QTreeWidgetItem&",
		classes["QTreeWidgetItem"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1049] = type_
	type_ = &Type{
		"QTreeWidgetItem*",
		classes["QTreeWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1050] = type_
	type_ = &Type{
		"QTreeWidgetItem::ChildIndicatorPolicy",
		classes["QTreeWidgetItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1051] = type_
	type_ = &Type{
		"QTreeWidgetItem::ItemType",
		classes["QTreeWidgetItem"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1052] = type_
	type_ = &Type{
		"QTreeWidgetItemIterator&",
		classes["QTreeWidgetItemIterator"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1053] = type_
	type_ = &Type{
		"QTreeWidgetItemIterator*",
		classes["QTreeWidgetItemIterator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1054] = type_
	type_ = &Type{
		"QTreeWidgetItemIterator::IteratorFlag",
		classes["QTreeWidgetItemIterator"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1055] = type_
	type_ = &Type{
		"QUndoCommand*",
		classes["QUndoCommand"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1056] = type_
	type_ = &Type{
		"QUndoGroup*",
		classes["QUndoGroup"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1057] = type_
	type_ = &Type{
		"QUndoStack*",
		classes["QUndoStack"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1058] = type_
	type_ = &Type{
		"QUndoView*",
		classes["QUndoView"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1059] = type_
	type_ = &Type{
		"QUnixPrintWidget*",
		classes["QUnixPrintWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1060] = type_
	type_ = &Type{
		"QUrl",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1061] = type_
	type_ = &Type{
		"QUrl&",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1062] = type_
	type_ = &Type{
		"QUrl::FormattingOption",
		classes["QUrl"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1063] = type_
	type_ = &Type{
		"QUuid&",
		classes["QUuid"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1064] = type_
	type_ = &Type{
		"QVBoxLayout*",
		classes["QVBoxLayout"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1065] = type_
	type_ = &Type{
		"QValidator*",
		classes["QValidator"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1066] = type_
	type_ = &Type{
		"QValidator::State",
		classes["QValidator"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1067] = type_
	type_ = &Type{
		"QVariant",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1068] = type_
	type_ = &Type{
		"QVariant&",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1069] = type_
	type_ = &Type{
		"QVariant::Type",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1070] = type_
	type_ = &Type{
		"QVariant::Type&",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1071] = type_
	type_ = &Type{
		"QVector2D",
		classes["QVector2D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1072] = type_
	type_ = &Type{
		"QVector2D&",
		classes["QVector2D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1073] = type_
	type_ = &Type{
		"QVector2D*",
		classes["QVector2D"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1074] = type_
	type_ = &Type{
		"QVector3D",
		classes["QVector3D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1075] = type_
	type_ = &Type{
		"QVector3D&",
		classes["QVector3D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1076] = type_
	type_ = &Type{
		"QVector3D*",
		classes["QVector3D"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1077] = type_
	type_ = &Type{
		"QVector4D",
		classes["QVector4D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1078] = type_
	type_ = &Type{
		"QVector4D&",
		classes["QVector4D"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1079] = type_
	type_ = &Type{
		"QVector4D*",
		classes["QVector4D"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1080] = type_
	type_ = &Type{
		"QVector<QAbstractTextDocumentLayout::Selection>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1081] = type_
	type_ = &Type{
		"QVector<QPair<double,QColor> >",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1082] = type_
	type_ = &Type{
		"QVector<QPointF>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1083] = type_
	type_ = &Type{
		"QVector<QRect>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1084] = type_
	type_ = &Type{
		"QVector<QTextFormat>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1085] = type_
	type_ = &Type{
		"QVector<QTextLength>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1086] = type_
	type_ = &Type{
		"QVector<double>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1087] = type_
	type_ = &Type{
		"QVector<unsigned int>",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1088] = type_
	type_ = &Type{
		"QWhatsThis*",
		classes["QWhatsThis"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1089] = type_
	type_ = &Type{
		"QWhatsThisClickedEvent*",
		classes["QWhatsThisClickedEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1090] = type_
	type_ = &Type{
		"QWheelEvent*",
		classes["QWheelEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1091] = type_
	type_ = &Type{
		"QWidget*",
		classes["QWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1092] = type_
	type_ = &Type{
		"QWidget::RenderFlag",
		classes["QWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1093] = type_
	type_ = &Type{
		"QWidgetAction*",
		classes["QWidgetAction"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1094] = type_
	type_ = &Type{
		"QWidgetItem*",
		classes["QWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1095] = type_
	type_ = &Type{
		"QWidgetItemV2*",
		classes["QWidgetItemV2"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1096] = type_
	type_ = &Type{
		"QWindowStateChangeEvent*",
		classes["QWindowStateChangeEvent"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1097] = type_
	type_ = &Type{
		"QWindowSurface*",
		classes["QWindowSurface"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1098] = type_
	type_ = &Type{
		"QWizard*",
		classes["QWizard"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1099] = type_
	type_ = &Type{
		"QWizard::WizardButton",
		classes["QWizard"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1100] = type_
	type_ = &Type{
		"QWizard::WizardOption",
		classes["QWizard"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1101] = type_
	type_ = &Type{
		"QWizard::WizardPixmap",
		classes["QWizard"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1102] = type_
	type_ = &Type{
		"QWizard::WizardStyle",
		classes["QWizard"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1103] = type_
	type_ = &Type{
		"QWizardPage*",
		classes["QWizardPage"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1104] = type_
	type_ = &Type{
		"QWorkspace*",
		classes["QWorkspace"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1105] = type_
	type_ = &Type{
		"QWorkspace::WindowOrder",
		classes["QWorkspace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1106] = type_
	type_ = &Type{
		"QX11EmbedContainer*",
		classes["QX11EmbedContainer"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1107] = type_
	type_ = &Type{
		"QX11EmbedContainer::Error",
		classes["QX11EmbedContainer"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1108] = type_
	type_ = &Type{
		"QX11EmbedWidget*",
		classes["QX11EmbedWidget"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1109] = type_
	type_ = &Type{
		"QX11EmbedWidget::Error",
		classes["QX11EmbedWidget"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1110] = type_
	type_ = &Type{
		"QX11Info&",
		classes["QX11Info"],
		T_CLASS,
		KIND_STACK,
		false,
	}
	types[1111] = type_
	type_ = &Type{
		"QX11Info*",
		classes["QX11Info"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1112] = type_
	type_ = &Type{
		"QX11InfoData*",
		classes["QX11InfoData"],
		T_CLASS,
		KIND_POINTER,
		false,
	}
	types[1113] = type_
	type_ = &Type{
		"Qt::Alignment",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[1114] = type_
	type_ = &Type{
		"Qt::AlignmentFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1115] = type_
	type_ = &Type{
		"Qt::AnchorAttribute",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1116] = type_
	type_ = &Type{
		"Qt::AnchorPoint",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1117] = type_
	type_ = &Type{
		"Qt::ApplicationAttribute",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1118] = type_
	type_ = &Type{
		"Qt::ArrowType",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1119] = type_
	type_ = &Type{
		"Qt::AspectRatioMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1120] = type_
	type_ = &Type{
		"Qt::Axis",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1121] = type_
	type_ = &Type{
		"Qt::BGMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1122] = type_
	type_ = &Type{
		"Qt::BrushStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1123] = type_
	type_ = &Type{
		"Qt::CaseSensitivity",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1124] = type_
	type_ = &Type{
		"Qt::CheckState",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1125] = type_
	type_ = &Type{
		"Qt::ClipOperation",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1126] = type_
	type_ = &Type{
		"Qt::ConnectionType",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1127] = type_
	type_ = &Type{
		"Qt::ContextMenuPolicy",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1128] = type_
	type_ = &Type{
		"Qt::CoordinateSystem",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1129] = type_
	type_ = &Type{
		"Qt::Corner",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1130] = type_
	type_ = &Type{
		"Qt::CursorMoveStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1131] = type_
	type_ = &Type{
		"Qt::CursorShape",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1132] = type_
	type_ = &Type{
		"Qt::DateFormat",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1133] = type_
	type_ = &Type{
		"Qt::DayOfWeek",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1134] = type_
	type_ = &Type{
		"Qt::DockWidgetArea",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1135] = type_
	type_ = &Type{
		"Qt::DockWidgetAreaSizes",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1136] = type_
	type_ = &Type{
		"Qt::DockWidgetAreas",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[1137] = type_
	type_ = &Type{
		"Qt::DropAction",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1138] = type_
	type_ = &Type{
		"Qt::EventPriority",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1139] = type_
	type_ = &Type{
		"Qt::FillRule",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1140] = type_
	type_ = &Type{
		"Qt::FocusPolicy",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1141] = type_
	type_ = &Type{
		"Qt::FocusReason",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1142] = type_
	type_ = &Type{
		"Qt::GestureFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1143] = type_
	type_ = &Type{
		"Qt::GestureState",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1144] = type_
	type_ = &Type{
		"Qt::GestureType",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1145] = type_
	type_ = &Type{
		"Qt::GlobalColor",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1146] = type_
	type_ = &Type{
		"Qt::HitTestAccuracy",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1147] = type_
	type_ = &Type{
		"Qt::ImageConversionFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1148] = type_
	type_ = &Type{
		"Qt::Initialization",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1149] = type_
	type_ = &Type{
		"Qt::InputMethodHint",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1150] = type_
	type_ = &Type{
		"Qt::InputMethodQuery",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1151] = type_
	type_ = &Type{
		"Qt::ItemDataRole",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1152] = type_
	type_ = &Type{
		"Qt::ItemFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1153] = type_
	type_ = &Type{
		"Qt::ItemSelectionMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1154] = type_
	type_ = &Type{
		"Qt::Key",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1155] = type_
	type_ = &Type{
		"Qt::KeyboardModifier",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1156] = type_
	type_ = &Type{
		"Qt::LayoutDirection",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1157] = type_
	type_ = &Type{
		"Qt::MaskMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1158] = type_
	type_ = &Type{
		"Qt::MatchFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1159] = type_
	type_ = &Type{
		"Qt::Modifier",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1160] = type_
	type_ = &Type{
		"Qt::MouseButton",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1161] = type_
	type_ = &Type{
		"Qt::NavigationMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1162] = type_
	type_ = &Type{
		"Qt::Orientation",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1163] = type_
	type_ = &Type{
		"Qt::PenCapStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1164] = type_
	type_ = &Type{
		"Qt::PenJoinStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1165] = type_
	type_ = &Type{
		"Qt::PenStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1166] = type_
	type_ = &Type{
		"Qt::ScrollBarPolicy",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1167] = type_
	type_ = &Type{
		"Qt::ShortcutContext",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1168] = type_
	type_ = &Type{
		"Qt::SizeHint",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1169] = type_
	type_ = &Type{
		"Qt::SizeMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1170] = type_
	type_ = &Type{
		"Qt::SortOrder",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1171] = type_
	type_ = &Type{
		"Qt::TextElideMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1172] = type_
	type_ = &Type{
		"Qt::TextFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1173] = type_
	type_ = &Type{
		"Qt::TextFormat",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1174] = type_
	type_ = &Type{
		"Qt::TextInteractionFlag",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1175] = type_
	type_ = &Type{
		"Qt::TileRule",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1176] = type_
	type_ = &Type{
		"Qt::TimeSpec",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1177] = type_
	type_ = &Type{
		"Qt::ToolBarArea",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1178] = type_
	type_ = &Type{
		"Qt::ToolBarAreaSizes",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1179] = type_
	type_ = &Type{
		"Qt::ToolBarAreas",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[1180] = type_
	type_ = &Type{
		"Qt::ToolButtonStyle",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1181] = type_
	type_ = &Type{
		"Qt::TouchPointState",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1182] = type_
	type_ = &Type{
		"Qt::TransformationMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1183] = type_
	type_ = &Type{
		"Qt::UIEffect",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1184] = type_
	type_ = &Type{
		"Qt::WhiteSpaceMode",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1185] = type_
	type_ = &Type{
		"Qt::WidgetAttribute",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1186] = type_
	type_ = &Type{
		"Qt::WindowFrameSection",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1187] = type_
	type_ = &Type{
		"Qt::WindowModality",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1188] = type_
	type_ = &Type{
		"Qt::WindowState",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1189] = type_
	type_ = &Type{
		"Qt::WindowStates",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[1190] = type_
	type_ = &Type{
		"Qt::WindowType",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1191] = type_
	type_ = &Type{
		"QtConcurrent::ReduceOption",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1192] = type_
	type_ = &Type{
		"QtConcurrent::ThreadFunctionResult",
		nil,
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1193] = type_
	type_ = &Type{
		"QtMsgType",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1194] = type_
	type_ = &Type{
		"QtValidLicenseForActiveQtModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1195] = type_
	type_ = &Type{
		"QtValidLicenseForCoreModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1196] = type_
	type_ = &Type{
		"QtValidLicenseForDBusModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1197] = type_
	type_ = &Type{
		"QtValidLicenseForDeclarativeModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1198] = type_
	type_ = &Type{
		"QtValidLicenseForGuiModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1199] = type_
	type_ = &Type{
		"QtValidLicenseForHelpModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1200] = type_
	type_ = &Type{
		"QtValidLicenseForMultimediaModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1201] = type_
	type_ = &Type{
		"QtValidLicenseForNetworkModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1202] = type_
	type_ = &Type{
		"QtValidLicenseForOpenGLModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1203] = type_
	type_ = &Type{
		"QtValidLicenseForOpenVGModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1204] = type_
	type_ = &Type{
		"QtValidLicenseForQt3SupportLightModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1205] = type_
	type_ = &Type{
		"QtValidLicenseForQt3SupportModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1206] = type_
	type_ = &Type{
		"QtValidLicenseForScriptModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1207] = type_
	type_ = &Type{
		"QtValidLicenseForScriptToolsModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1208] = type_
	type_ = &Type{
		"QtValidLicenseForSqlModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1209] = type_
	type_ = &Type{
		"QtValidLicenseForSvgModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1210] = type_
	type_ = &Type{
		"QtValidLicenseForTestModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1211] = type_
	type_ = &Type{
		"QtValidLicenseForXmlModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1212] = type_
	type_ = &Type{
		"QtValidLicenseForXmlPatternsModule",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1213] = type_
	type_ = &Type{
		"_XDisplay*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1214] = type_
	type_ = &Type{
		"_XEvent*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1215] = type_
	type_ = &Type{
		"_XRegion*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1216] = type_
	type_ = &Type{
		"__codecvt_result",
		classes["QGlobalSpace"],
		T_ENUM,
		KIND_STACK,
		false,
	}
	types[1217] = type_
	type_ = &Type{
		"bool",
		nil,
		T_BOOL,
		KIND_STACK,
		false,
	}
	types[1218] = type_
	type_ = &Type{
		"bool*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1219] = type_
	type_ = &Type{
		"char",
		nil,
		T_CHAR,
		KIND_STACK,
		false,
	}
	types[1220] = type_
	type_ = &Type{
		"char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1221] = type_
	type_ = &Type{
		"char**",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1222] = type_
	type_ = &Type{
		"const QAbstractItemModel*",
		classes["QAbstractItemModel"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1223] = type_
	type_ = &Type{
		"const QAbstractTextDocumentLayout::PaintContext&",
		classes["QAbstractTextDocumentLayout::PaintContext"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1224] = type_
	type_ = &Type{
		"const QAbstractTextDocumentLayout::Selection&",
		classes["QAbstractTextDocumentLayout::Selection"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1225] = type_
	type_ = &Type{
		"const QAbstractUndoItem&",
		classes["QAbstractUndoItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1226] = type_
	type_ = &Type{
		"const QAccessible&",
		classes["QAccessible"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1227] = type_
	type_ = &Type{
		"const QAccessible2::TableModelChange&",
		classes["QAccessible2::TableModelChange"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1228] = type_
	type_ = &Type{
		"const QAccessible2Interface&",
		classes["QAccessible2Interface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1229] = type_
	type_ = &Type{
		"const QAccessibleActionInterface&",
		classes["QAccessibleActionInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1230] = type_
	type_ = &Type{
		"const QAccessibleBridge&",
		classes["QAccessibleBridge"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1231] = type_
	type_ = &Type{
		"const QAccessibleBridgeFactoryInterface&",
		classes["QAccessibleBridgeFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1232] = type_
	type_ = &Type{
		"const QAccessibleEditableTextInterface&",
		classes["QAccessibleEditableTextInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1233] = type_
	type_ = &Type{
		"const QAccessibleEvent&",
		classes["QAccessibleEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1234] = type_
	type_ = &Type{
		"const QAccessibleFactoryInterface&",
		classes["QAccessibleFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1235] = type_
	type_ = &Type{
		"const QAccessibleImageInterface&",
		classes["QAccessibleImageInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1236] = type_
	type_ = &Type{
		"const QAccessibleInterface&",
		classes["QAccessibleInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1237] = type_
	type_ = &Type{
		"const QAccessibleInterface*",
		classes["QAccessibleInterface"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1238] = type_
	type_ = &Type{
		"const QAccessibleInterfaceEx&",
		classes["QAccessibleInterfaceEx"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1239] = type_
	type_ = &Type{
		"const QAccessibleSimpleEditableTextInterface&",
		classes["QAccessibleSimpleEditableTextInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1240] = type_
	type_ = &Type{
		"const QAccessibleTable2CellInterface&",
		classes["QAccessibleTable2CellInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1241] = type_
	type_ = &Type{
		"const QAccessibleTable2Interface&",
		classes["QAccessibleTable2Interface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1242] = type_
	type_ = &Type{
		"const QAccessibleTableInterface&",
		classes["QAccessibleTableInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1243] = type_
	type_ = &Type{
		"const QAccessibleTextInterface&",
		classes["QAccessibleTextInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1244] = type_
	type_ = &Type{
		"const QAccessibleValueInterface&",
		classes["QAccessibleValueInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1245] = type_
	type_ = &Type{
		"const QAction*",
		classes["QAction"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1246] = type_
	type_ = &Type{
		"const QActionEvent&",
		classes["QActionEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1247] = type_
	type_ = &Type{
		"const QBitArray&",
		classes["QBitArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1248] = type_
	type_ = &Type{
		"const QBitmap&",
		classes["QBitmap"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1249] = type_
	type_ = &Type{
		"const QBitmap*",
		classes["QBitmap"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1250] = type_
	type_ = &Type{
		"const QBrush&",
		classes["QBrush"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1251] = type_
	type_ = &Type{
		"const QBrush*",
		classes["QBrush"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1252] = type_
	type_ = &Type{
		"const QByteArray",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1253] = type_
	type_ = &Type{
		"const QByteArray&",
		classes["QByteArray"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1254] = type_
	type_ = &Type{
		"const QChar&",
		classes["QChar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1255] = type_
	type_ = &Type{
		"const QChar*",
		classes["QChar"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1256] = type_
	type_ = &Type{
		"const QClipboardEvent&",
		classes["QClipboardEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1257] = type_
	type_ = &Type{
		"const QCloseEvent&",
		classes["QCloseEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1258] = type_
	type_ = &Type{
		"const QColor",
		classes["QColor"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1259] = type_
	type_ = &Type{
		"const QColor&",
		classes["QColor"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1260] = type_
	type_ = &Type{
		"const QColormap&",
		classes["QColormap"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1261] = type_
	type_ = &Type{
		"const QConicalGradient&",
		classes["QConicalGradient"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1262] = type_
	type_ = &Type{
		"const QContextMenuEvent&",
		classes["QContextMenuEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1263] = type_
	type_ = &Type{
		"const QCursor&",
		classes["QCursor"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1264] = type_
	type_ = &Type{
		"const QDate&",
		classes["QDate"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1265] = type_
	type_ = &Type{
		"const QDateTime&",
		classes["QDateTime"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1266] = type_
	type_ = &Type{
		"const QDesktopServices&",
		classes["QDesktopServices"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1267] = type_
	type_ = &Type{
		"const QDir&",
		classes["QDir"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1268] = type_
	type_ = &Type{
		"const QDragEnterEvent&",
		classes["QDragEnterEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1269] = type_
	type_ = &Type{
		"const QDragLeaveEvent&",
		classes["QDragLeaveEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1270] = type_
	type_ = &Type{
		"const QDragMoveEvent&",
		classes["QDragMoveEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1271] = type_
	type_ = &Type{
		"const QDragResponseEvent&",
		classes["QDragResponseEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1272] = type_
	type_ = &Type{
		"const QDropEvent&",
		classes["QDropEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1273] = type_
	type_ = &Type{
		"const QEasingCurve&",
		classes["QEasingCurve"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1274] = type_
	type_ = &Type{
		"const QEvent*",
		classes["QEvent"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1275] = type_
	type_ = &Type{
		"const QFileDialogArgs&",
		classes["QFileDialogArgs"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1276] = type_
	type_ = &Type{
		"const QFileInfo&",
		classes["QFileInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1277] = type_
	type_ = &Type{
		"const QFileOpenEvent&",
		classes["QFileOpenEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1278] = type_
	type_ = &Type{
		"const QFocusEvent&",
		classes["QFocusEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1279] = type_
	type_ = &Type{
		"const QFont&",
		classes["QFont"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1280] = type_
	type_ = &Type{
		"const QFontDatabase&",
		classes["QFontDatabase"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1281] = type_
	type_ = &Type{
		"const QFontInfo&",
		classes["QFontInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1282] = type_
	type_ = &Type{
		"const QFontMetrics&",
		classes["QFontMetrics"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1283] = type_
	type_ = &Type{
		"const QFontMetricsF&",
		classes["QFontMetricsF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1284] = type_
	type_ = &Type{
		"const QGestureEvent&",
		classes["QGestureEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1285] = type_
	type_ = &Type{
		"const QGestureRecognizer&",
		classes["QGestureRecognizer"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1286] = type_
	type_ = &Type{
		"const QGlyphRun&",
		classes["QGlyphRun"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1287] = type_
	type_ = &Type{
		"const QGradient&",
		classes["QGradient"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1288] = type_
	type_ = &Type{
		"const QGradient*",
		classes["QGradient"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1289] = type_
	type_ = &Type{
		"const QGraphicsItem*",
		classes["QGraphicsItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1290] = type_
	type_ = &Type{
		"const QGraphicsObject*",
		classes["QGraphicsObject"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1291] = type_
	type_ = &Type{
		"const QGraphicsSceneEventPrivate*",
		classes["QGraphicsSceneEventPrivate"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1292] = type_
	type_ = &Type{
		"const QHashDummyValue&",
		classes["QHashDummyValue"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1293] = type_
	type_ = &Type{
		"const QHelpEvent&",
		classes["QHelpEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1294] = type_
	type_ = &Type{
		"const QHideEvent&",
		classes["QHideEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1295] = type_
	type_ = &Type{
		"const QHoverEvent&",
		classes["QHoverEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1296] = type_
	type_ = &Type{
		"const QIcon&",
		classes["QIcon"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1297] = type_
	type_ = &Type{
		"const QIconDragEvent&",
		classes["QIconDragEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1298] = type_
	type_ = &Type{
		"const QIconEngine&",
		classes["QIconEngine"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1299] = type_
	type_ = &Type{
		"const QIconEngineFactoryInterface&",
		classes["QIconEngineFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1300] = type_
	type_ = &Type{
		"const QIconEngineFactoryInterfaceV2&",
		classes["QIconEngineFactoryInterfaceV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1301] = type_
	type_ = &Type{
		"const QIconEngineV2&",
		classes["QIconEngineV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1302] = type_
	type_ = &Type{
		"const QIconEngineV2::AvailableSizesArgument&",
		classes["QIconEngineV2::AvailableSizesArgument"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1303] = type_
	type_ = &Type{
		"const QImage&",
		classes["QImage"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1304] = type_
	type_ = &Type{
		"const QImageIOHandlerFactoryInterface&",
		classes["QImageIOHandlerFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1305] = type_
	type_ = &Type{
		"const QImageTextKeyLang&",
		classes["QImageTextKeyLang"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1306] = type_
	type_ = &Type{
		"const QInputContextFactory&",
		classes["QInputContextFactory"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1307] = type_
	type_ = &Type{
		"const QInputContextFactoryInterface&",
		classes["QInputContextFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1308] = type_
	type_ = &Type{
		"const QInputEvent&",
		classes["QInputEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1309] = type_
	type_ = &Type{
		"const QInputMethodEvent&",
		classes["QInputMethodEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1310] = type_
	type_ = &Type{
		"const QInputMethodEvent::Attribute&",
		classes["QInputMethodEvent::Attribute"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1311] = type_
	type_ = &Type{
		"const QItemEditorCreatorBase&",
		classes["QItemEditorCreatorBase"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1312] = type_
	type_ = &Type{
		"const QItemEditorFactory&",
		classes["QItemEditorFactory"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1313] = type_
	type_ = &Type{
		"const QItemEditorFactory*",
		classes["QItemEditorFactory"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1314] = type_
	type_ = &Type{
		"const QItemSelection",
		classes["QItemSelection"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1315] = type_
	type_ = &Type{
		"const QItemSelection&",
		classes["QItemSelection"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1316] = type_
	type_ = &Type{
		"const QItemSelectionRange&",
		classes["QItemSelectionRange"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1317] = type_
	type_ = &Type{
		"const QKeyEvent&",
		classes["QKeyEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1318] = type_
	type_ = &Type{
		"const QKeySequence&",
		classes["QKeySequence"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1319] = type_
	type_ = &Type{
		"const QLatin1String&",
		classes["QLatin1String"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1320] = type_
	type_ = &Type{
		"const QLayoutItem&",
		classes["QLayoutItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1321] = type_
	type_ = &Type{
		"const QLine&",
		classes["QLine"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1322] = type_
	type_ = &Type{
		"const QLine*",
		classes["QLine"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1323] = type_
	type_ = &Type{
		"const QLineF&",
		classes["QLineF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1324] = type_
	type_ = &Type{
		"const QLineF*",
		classes["QLineF"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1325] = type_
	type_ = &Type{
		"const QLinearGradient&",
		classes["QLinearGradient"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1326] = type_
	type_ = &Type{
		"const QList<QByteArray>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1327] = type_
	type_ = &Type{
		"const QList<QGesture*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1328] = type_
	type_ = &Type{
		"const QList<QGraphicsItem*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1329] = type_
	type_ = &Type{
		"const QList<QGraphicsTransform*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1330] = type_
	type_ = &Type{
		"const QList<QInputMethodEvent::Attribute>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1331] = type_
	type_ = &Type{
		"const QList<QKeySequence>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1332] = type_
	type_ = &Type{
		"const QList<QListWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1333] = type_
	type_ = &Type{
		"const QList<QModelIndex>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1334] = type_
	type_ = &Type{
		"const QList<QObject*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1335] = type_
	type_ = &Type{
		"const QList<QRectF>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1336] = type_
	type_ = &Type{
		"const QList<QSize>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1337] = type_
	type_ = &Type{
		"const QList<QStandardItem*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1338] = type_
	type_ = &Type{
		"const QList<QTableWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1339] = type_
	type_ = &Type{
		"const QList<QTextEdit::ExtraSelection>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1340] = type_
	type_ = &Type{
		"const QList<QTextLayout::FormatRange>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1341] = type_
	type_ = &Type{
		"const QList<QTextOption::Tab>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1342] = type_
	type_ = &Type{
		"const QList<QTouchEvent::TouchPoint>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1343] = type_
	type_ = &Type{
		"const QList<QTreeWidgetItem*>",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1344] = type_
	type_ = &Type{
		"const QList<QTreeWidgetItem*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1345] = type_
	type_ = &Type{
		"const QList<QUrl>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1346] = type_
	type_ = &Type{
		"const QList<QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1347] = type_
	type_ = &Type{
		"const QList<QWidget*>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1348] = type_
	type_ = &Type{
		"const QList<QWizard::WizardButton>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1349] = type_
	type_ = &Type{
		"const QList<int>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1350] = type_
	type_ = &Type{
		"const QListWidgetItem&",
		classes["QListWidgetItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1351] = type_
	type_ = &Type{
		"const QListWidgetItem*",
		classes["QListWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1352] = type_
	type_ = &Type{
		"const QLocale&",
		classes["QLocale"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1353] = type_
	type_ = &Type{
		"const QMap<int,QVariant>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1354] = type_
	type_ = &Type{
		"const QMargins&",
		classes["QMargins"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1355] = type_
	type_ = &Type{
		"const QMatrix&",
		classes["QMatrix"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1356] = type_
	type_ = &Type{
		"const QMatrix4x4&",
		classes["QMatrix4x4"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1357] = type_
	type_ = &Type{
		"const QMetaObject&",
		classes["QMetaObject"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1358] = type_
	type_ = &Type{
		"const QMetaObject*",
		classes["QMetaObject"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1359] = type_
	type_ = &Type{
		"const QMimeData*",
		classes["QMimeData"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1360] = type_
	type_ = &Type{
		"const QMimeSource&",
		classes["QMimeSource"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1361] = type_
	type_ = &Type{
		"const QModelIndex&",
		classes["QModelIndex"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1362] = type_
	type_ = &Type{
		"const QModelIndexList&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1363] = type_
	type_ = &Type{
		"const QMouseEvent&",
		classes["QMouseEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1364] = type_
	type_ = &Type{
		"const QMoveEvent&",
		classes["QMoveEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1365] = type_
	type_ = &Type{
		"const QObject*",
		classes["QObject"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1366] = type_
	type_ = &Type{
		"const QPaintDevice*",
		classes["QPaintDevice"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1367] = type_
	type_ = &Type{
		"const QPaintEngineState&",
		classes["QPaintEngineState"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1368] = type_
	type_ = &Type{
		"const QPaintEvent&",
		classes["QPaintEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1369] = type_
	type_ = &Type{
		"const QPainter::PixmapFragment&",
		classes["QPainter::PixmapFragment"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1370] = type_
	type_ = &Type{
		"const QPainter::PixmapFragment*",
		classes["QPainter::PixmapFragment"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1371] = type_
	type_ = &Type{
		"const QPainterPath&",
		classes["QPainterPath"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1372] = type_
	type_ = &Type{
		"const QPainterPath::Element&",
		classes["QPainterPath::Element"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1373] = type_
	type_ = &Type{
		"const QPalette&",
		classes["QPalette"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1374] = type_
	type_ = &Type{
		"const QPen&",
		classes["QPen"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1375] = type_
	type_ = &Type{
		"const QPersistentModelIndex&",
		classes["QPersistentModelIndex"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1376] = type_
	type_ = &Type{
		"const QPicture&",
		classes["QPicture"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1377] = type_
	type_ = &Type{
		"const QPicture*",
		classes["QPicture"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1378] = type_
	type_ = &Type{
		"const QPictureFormatInterface&",
		classes["QPictureFormatInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1379] = type_
	type_ = &Type{
		"const QPixmap",
		classes["QPixmap"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1380] = type_
	type_ = &Type{
		"const QPixmap&",
		classes["QPixmap"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1381] = type_
	type_ = &Type{
		"const QPixmap*",
		classes["QPixmap"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1382] = type_
	type_ = &Type{
		"const QPixmapCache&",
		classes["QPixmapCache"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1383] = type_
	type_ = &Type{
		"const QPixmapCache::Key&",
		classes["QPixmapCache::Key"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1384] = type_
	type_ = &Type{
		"const QPoint",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1385] = type_
	type_ = &Type{
		"const QPoint&",
		classes["QPoint"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1386] = type_
	type_ = &Type{
		"const QPoint*",
		classes["QPoint"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1387] = type_
	type_ = &Type{
		"const QPointF",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1388] = type_
	type_ = &Type{
		"const QPointF&",
		classes["QPointF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1389] = type_
	type_ = &Type{
		"const QPointF*",
		classes["QPointF"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1390] = type_
	type_ = &Type{
		"const QPolygon&",
		classes["QPolygon"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1391] = type_
	type_ = &Type{
		"const QPolygonF&",
		classes["QPolygonF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1392] = type_
	type_ = &Type{
		"const QPrintEngine&",
		classes["QPrintEngine"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1393] = type_
	type_ = &Type{
		"const QPrinter&",
		classes["QPrinter"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1394] = type_
	type_ = &Type{
		"const QPrinterInfo&",
		classes["QPrinterInfo"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1395] = type_
	type_ = &Type{
		"const QQuaternion",
		classes["QQuaternion"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1396] = type_
	type_ = &Type{
		"const QQuaternion&",
		classes["QQuaternion"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1397] = type_
	type_ = &Type{
		"const QRadialGradient&",
		classes["QRadialGradient"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1398] = type_
	type_ = &Type{
		"const QRawFont&",
		classes["QRawFont"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1399] = type_
	type_ = &Type{
		"const QRect",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1400] = type_
	type_ = &Type{
		"const QRect&",
		classes["QRect"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1401] = type_
	type_ = &Type{
		"const QRect*",
		classes["QRect"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1402] = type_
	type_ = &Type{
		"const QRectF&",
		classes["QRectF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1403] = type_
	type_ = &Type{
		"const QRectF*",
		classes["QRectF"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1404] = type_
	type_ = &Type{
		"const QRegExp&",
		classes["QRegExp"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1405] = type_
	type_ = &Type{
		"const QRegExp*",
		classes["QRegExp"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1406] = type_
	type_ = &Type{
		"const QRegion",
		classes["QRegion"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1407] = type_
	type_ = &Type{
		"const QRegion&",
		classes["QRegion"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1408] = type_
	type_ = &Type{
		"const QResizeEvent&",
		classes["QResizeEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1409] = type_
	type_ = &Type{
		"const QShortcutEvent&",
		classes["QShortcutEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1410] = type_
	type_ = &Type{
		"const QShowEvent&",
		classes["QShowEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1411] = type_
	type_ = &Type{
		"const QSize",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1412] = type_
	type_ = &Type{
		"const QSize&",
		classes["QSize"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1413] = type_
	type_ = &Type{
		"const QSizeF",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1414] = type_
	type_ = &Type{
		"const QSizeF&",
		classes["QSizeF"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1415] = type_
	type_ = &Type{
		"const QSizePolicy&",
		classes["QSizePolicy"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1416] = type_
	type_ = &Type{
		"const QSpacerItem&",
		classes["QSpacerItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1417] = type_
	type_ = &Type{
		"const QSplitter&",
		classes["QSplitter"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1418] = type_
	type_ = &Type{
		"const QStandardItem&",
		classes["QStandardItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1419] = type_
	type_ = &Type{
		"const QStandardItem*",
		classes["QStandardItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1420] = type_
	type_ = &Type{
		"const QStaticText&",
		classes["QStaticText"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1421] = type_
	type_ = &Type{
		"const QStatusTipEvent&",
		classes["QStatusTipEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1422] = type_
	type_ = &Type{
		"const QString",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1423] = type_
	type_ = &Type{
		"const QString&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1424] = type_
	type_ = &Type{
		"const QStringList&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1425] = type_
	type_ = &Type{
		"const QStringList*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1426] = type_
	type_ = &Type{
		"const QStringRef&",
		classes["QStringRef"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1427] = type_
	type_ = &Type{
		"const QStyle*",
		classes["QStyle"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1428] = type_
	type_ = &Type{
		"const QStyleFactory&",
		classes["QStyleFactory"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1429] = type_
	type_ = &Type{
		"const QStyleFactoryInterface&",
		classes["QStyleFactoryInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1430] = type_
	type_ = &Type{
		"const QStyleHintReturn&",
		classes["QStyleHintReturn"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1431] = type_
	type_ = &Type{
		"const QStyleHintReturnMask&",
		classes["QStyleHintReturnMask"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1432] = type_
	type_ = &Type{
		"const QStyleHintReturnVariant&",
		classes["QStyleHintReturnVariant"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1433] = type_
	type_ = &Type{
		"const QStyleOption&",
		classes["QStyleOption"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1434] = type_
	type_ = &Type{
		"const QStyleOption*",
		classes["QStyleOption"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1435] = type_
	type_ = &Type{
		"const QStyleOption::OptionType&",
		classes["QStyleOption"],
		T_ENUM,
		KIND_STACK,
		true,
	}
	types[1436] = type_
	type_ = &Type{
		"const QStyleOptionButton&",
		classes["QStyleOptionButton"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1437] = type_
	type_ = &Type{
		"const QStyleOptionComboBox&",
		classes["QStyleOptionComboBox"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1438] = type_
	type_ = &Type{
		"const QStyleOptionComplex&",
		classes["QStyleOptionComplex"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1439] = type_
	type_ = &Type{
		"const QStyleOptionComplex*",
		classes["QStyleOptionComplex"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1440] = type_
	type_ = &Type{
		"const QStyleOptionDockWidget&",
		classes["QStyleOptionDockWidget"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1441] = type_
	type_ = &Type{
		"const QStyleOptionDockWidgetV2&",
		classes["QStyleOptionDockWidgetV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1442] = type_
	type_ = &Type{
		"const QStyleOptionFocusRect&",
		classes["QStyleOptionFocusRect"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1443] = type_
	type_ = &Type{
		"const QStyleOptionFrame&",
		classes["QStyleOptionFrame"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1444] = type_
	type_ = &Type{
		"const QStyleOptionFrameV2&",
		classes["QStyleOptionFrameV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1445] = type_
	type_ = &Type{
		"const QStyleOptionFrameV3&",
		classes["QStyleOptionFrameV3"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1446] = type_
	type_ = &Type{
		"const QStyleOptionGraphicsItem&",
		classes["QStyleOptionGraphicsItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1447] = type_
	type_ = &Type{
		"const QStyleOptionGraphicsItem*",
		classes["QStyleOptionGraphicsItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1448] = type_
	type_ = &Type{
		"const QStyleOptionGroupBox&",
		classes["QStyleOptionGroupBox"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1449] = type_
	type_ = &Type{
		"const QStyleOptionHeader&",
		classes["QStyleOptionHeader"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1450] = type_
	type_ = &Type{
		"const QStyleOptionMenuItem&",
		classes["QStyleOptionMenuItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1451] = type_
	type_ = &Type{
		"const QStyleOptionProgressBar&",
		classes["QStyleOptionProgressBar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1452] = type_
	type_ = &Type{
		"const QStyleOptionProgressBarV2&",
		classes["QStyleOptionProgressBarV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1453] = type_
	type_ = &Type{
		"const QStyleOptionRubberBand&",
		classes["QStyleOptionRubberBand"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1454] = type_
	type_ = &Type{
		"const QStyleOptionSizeGrip&",
		classes["QStyleOptionSizeGrip"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1455] = type_
	type_ = &Type{
		"const QStyleOptionSlider&",
		classes["QStyleOptionSlider"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1456] = type_
	type_ = &Type{
		"const QStyleOptionSpinBox&",
		classes["QStyleOptionSpinBox"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1457] = type_
	type_ = &Type{
		"const QStyleOptionTab&",
		classes["QStyleOptionTab"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1458] = type_
	type_ = &Type{
		"const QStyleOptionTabBarBase&",
		classes["QStyleOptionTabBarBase"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1459] = type_
	type_ = &Type{
		"const QStyleOptionTabBarBaseV2&",
		classes["QStyleOptionTabBarBaseV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1460] = type_
	type_ = &Type{
		"const QStyleOptionTabV2&",
		classes["QStyleOptionTabV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1461] = type_
	type_ = &Type{
		"const QStyleOptionTabV3&",
		classes["QStyleOptionTabV3"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1462] = type_
	type_ = &Type{
		"const QStyleOptionTabWidgetFrame&",
		classes["QStyleOptionTabWidgetFrame"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1463] = type_
	type_ = &Type{
		"const QStyleOptionTabWidgetFrameV2&",
		classes["QStyleOptionTabWidgetFrameV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1464] = type_
	type_ = &Type{
		"const QStyleOptionTitleBar&",
		classes["QStyleOptionTitleBar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1465] = type_
	type_ = &Type{
		"const QStyleOptionToolBar&",
		classes["QStyleOptionToolBar"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1466] = type_
	type_ = &Type{
		"const QStyleOptionToolBox&",
		classes["QStyleOptionToolBox"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1467] = type_
	type_ = &Type{
		"const QStyleOptionToolBoxV2&",
		classes["QStyleOptionToolBoxV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1468] = type_
	type_ = &Type{
		"const QStyleOptionToolButton&",
		classes["QStyleOptionToolButton"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1469] = type_
	type_ = &Type{
		"const QStyleOptionViewItem&",
		classes["QStyleOptionViewItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1470] = type_
	type_ = &Type{
		"const QStyleOptionViewItemV2&",
		classes["QStyleOptionViewItemV2"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1471] = type_
	type_ = &Type{
		"const QStyleOptionViewItemV3&",
		classes["QStyleOptionViewItemV3"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1472] = type_
	type_ = &Type{
		"const QStyleOptionViewItemV4&",
		classes["QStyleOptionViewItemV4"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1473] = type_
	type_ = &Type{
		"const QTableWidgetItem&",
		classes["QTableWidgetItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1474] = type_
	type_ = &Type{
		"const QTableWidgetItem*",
		classes["QTableWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1475] = type_
	type_ = &Type{
		"const QTableWidgetSelectionRange&",
		classes["QTableWidgetSelectionRange"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1476] = type_
	type_ = &Type{
		"const QTabletEvent&",
		classes["QTabletEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1477] = type_
	type_ = &Type{
		"const QTextBlock&",
		classes["QTextBlock"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1478] = type_
	type_ = &Type{
		"const QTextBlock::iterator&",
		classes["QTextBlock::iterator"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1479] = type_
	type_ = &Type{
		"const QTextBlockFormat&",
		classes["QTextBlockFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1480] = type_
	type_ = &Type{
		"const QTextBlockUserData&",
		classes["QTextBlockUserData"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1481] = type_
	type_ = &Type{
		"const QTextCharFormat&",
		classes["QTextCharFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1482] = type_
	type_ = &Type{
		"const QTextCursor&",
		classes["QTextCursor"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1483] = type_
	type_ = &Type{
		"const QTextDocument*",
		classes["QTextDocument"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1484] = type_
	type_ = &Type{
		"const QTextDocumentFragment&",
		classes["QTextDocumentFragment"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1485] = type_
	type_ = &Type{
		"const QTextEdit::ExtraSelection&",
		classes["QTextEdit::ExtraSelection"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1486] = type_
	type_ = &Type{
		"const QTextFormat&",
		classes["QTextFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1487] = type_
	type_ = &Type{
		"const QTextFragment&",
		classes["QTextFragment"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1488] = type_
	type_ = &Type{
		"const QTextFrame::iterator&",
		classes["QTextFrame::iterator"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1489] = type_
	type_ = &Type{
		"const QTextFrameFormat&",
		classes["QTextFrameFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1490] = type_
	type_ = &Type{
		"const QTextFrameLayoutData&",
		classes["QTextFrameLayoutData"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1491] = type_
	type_ = &Type{
		"const QTextImageFormat&",
		classes["QTextImageFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1492] = type_
	type_ = &Type{
		"const QTextInlineObject&",
		classes["QTextInlineObject"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1493] = type_
	type_ = &Type{
		"const QTextItem&",
		classes["QTextItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1494] = type_
	type_ = &Type{
		"const QTextLayout::FormatRange&",
		classes["QTextLayout::FormatRange"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1495] = type_
	type_ = &Type{
		"const QTextLayout::FormatRange*",
		classes["QTextLayout::FormatRange"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1496] = type_
	type_ = &Type{
		"const QTextLength&",
		classes["QTextLength"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1497] = type_
	type_ = &Type{
		"const QTextLine&",
		classes["QTextLine"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1498] = type_
	type_ = &Type{
		"const QTextListFormat&",
		classes["QTextListFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1499] = type_
	type_ = &Type{
		"const QTextObjectInterface&",
		classes["QTextObjectInterface"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1500] = type_
	type_ = &Type{
		"const QTextOption&",
		classes["QTextOption"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1501] = type_
	type_ = &Type{
		"const QTextOption::Tab&",
		classes["QTextOption::Tab"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1502] = type_
	type_ = &Type{
		"const QTextTableCell&",
		classes["QTextTableCell"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1503] = type_
	type_ = &Type{
		"const QTextTableCellFormat&",
		classes["QTextTableCellFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1504] = type_
	type_ = &Type{
		"const QTextTableFormat&",
		classes["QTextTableFormat"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1505] = type_
	type_ = &Type{
		"const QTileRules&",
		classes["QTileRules"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1506] = type_
	type_ = &Type{
		"const QTime&",
		classes["QTime"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1507] = type_
	type_ = &Type{
		"const QToolBarChangeEvent&",
		classes["QToolBarChangeEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1508] = type_
	type_ = &Type{
		"const QToolTip&",
		classes["QToolTip"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1509] = type_
	type_ = &Type{
		"const QTouchEvent&",
		classes["QTouchEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1510] = type_
	type_ = &Type{
		"const QTouchEvent::TouchPoint&",
		classes["QTouchEvent::TouchPoint"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1511] = type_
	type_ = &Type{
		"const QTransform&",
		classes["QTransform"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1512] = type_
	type_ = &Type{
		"const QTreeWidgetItem&",
		classes["QTreeWidgetItem"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1513] = type_
	type_ = &Type{
		"const QTreeWidgetItem*",
		classes["QTreeWidgetItem"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1514] = type_
	type_ = &Type{
		"const QTreeWidgetItemIterator",
		classes["QTreeWidgetItemIterator"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1515] = type_
	type_ = &Type{
		"const QTreeWidgetItemIterator&",
		classes["QTreeWidgetItemIterator"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1516] = type_
	type_ = &Type{
		"const QUndoCommand*",
		classes["QUndoCommand"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1517] = type_
	type_ = &Type{
		"const QUrl&",
		classes["QUrl"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1518] = type_
	type_ = &Type{
		"const QUuid&",
		classes["QUuid"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1519] = type_
	type_ = &Type{
		"const QValidator*",
		classes["QValidator"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1520] = type_
	type_ = &Type{
		"const QVariant&",
		classes["QVariant"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1521] = type_
	type_ = &Type{
		"const QVariant::Type",
		classes["QVariant"],
		T_ENUM,
		KIND_STACK,
		true,
	}
	types[1522] = type_
	type_ = &Type{
		"const QVariantComparisonHelper&",
		classes["QVariantComparisonHelper"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1523] = type_
	type_ = &Type{
		"const QVector2D",
		classes["QVector2D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1524] = type_
	type_ = &Type{
		"const QVector2D&",
		classes["QVector2D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1525] = type_
	type_ = &Type{
		"const QVector3D",
		classes["QVector3D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1526] = type_
	type_ = &Type{
		"const QVector3D&",
		classes["QVector3D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1527] = type_
	type_ = &Type{
		"const QVector4D",
		classes["QVector4D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1528] = type_
	type_ = &Type{
		"const QVector4D&",
		classes["QVector4D"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1529] = type_
	type_ = &Type{
		"const QVector<QAbstractTextDocumentLayout::Selection>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1530] = type_
	type_ = &Type{
		"const QVector<QColor>",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1531] = type_
	type_ = &Type{
		"const QVector<QLine>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1532] = type_
	type_ = &Type{
		"const QVector<QLineF>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1533] = type_
	type_ = &Type{
		"const QVector<QPair<double,QColor> >&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1534] = type_
	type_ = &Type{
		"const QVector<QPoint>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1535] = type_
	type_ = &Type{
		"const QVector<QPointF>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1536] = type_
	type_ = &Type{
		"const QVector<QRect>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1537] = type_
	type_ = &Type{
		"const QVector<QRectF>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1538] = type_
	type_ = &Type{
		"const QVector<QTextLayout::FormatRange>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1539] = type_
	type_ = &Type{
		"const QVector<QTextLength>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1540] = type_
	type_ = &Type{
		"const QVector<double>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1541] = type_
	type_ = &Type{
		"const QVector<unsigned int>",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1542] = type_
	type_ = &Type{
		"const QVector<unsigned int>&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1543] = type_
	type_ = &Type{
		"const QWhatsThis&",
		classes["QWhatsThis"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1544] = type_
	type_ = &Type{
		"const QWhatsThisClickedEvent&",
		classes["QWhatsThisClickedEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1545] = type_
	type_ = &Type{
		"const QWheelEvent&",
		classes["QWheelEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1546] = type_
	type_ = &Type{
		"const QWidget*",
		classes["QWidget"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1547] = type_
	type_ = &Type{
		"const QWindowStateChangeEvent&",
		classes["QWindowStateChangeEvent"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1548] = type_
	type_ = &Type{
		"const QX11Info&",
		classes["QX11Info"],
		T_CLASS,
		KIND_STACK,
		true,
	}
	types[1549] = type_
	type_ = &Type{
		"const QX11InfoData*",
		classes["QX11InfoData"],
		T_CLASS,
		KIND_POINTER,
		true,
	}
	types[1550] = type_
	type_ = &Type{
		"const bool",
		nil,
		T_BOOL,
		KIND_STACK,
		true,
	}
	types[1551] = type_
	type_ = &Type{
		"const char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1552] = type_
	type_ = &Type{
		"const char* const *",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1553] = type_
	type_ = &Type{
		"const double&",
		nil,
		T_VOIDP,
		KIND_STACK,
		true,
	}
	types[1554] = type_
	type_ = &Type{
		"const double*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1555] = type_
	type_ = &Type{
		"const int*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1556] = type_
	type_ = &Type{
		"const unsigned char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1557] = type_
	type_ = &Type{
		"const unsigned int*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1558] = type_
	type_ = &Type{
		"const void*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		true,
	}
	types[1559] = type_
	type_ = &Type{
		"double",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[1560] = type_
	type_ = &Type{
		"double&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1561] = type_
	type_ = &Type{
		"double*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1562] = type_
	type_ = &Type{
		"float",
		nil,
		T_FLOAT,
		KIND_STACK,
		false,
	}
	types[1563] = type_
	type_ = &Type{
		"int",
		nil,
		T_INT,
		KIND_STACK,
		false,
	}
	types[1564] = type_
	type_ = &Type{
		"int&",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1565] = type_
	type_ = &Type{
		"int(*)(const void*,const void*)",
		nil,
		T_INT,
		KIND_STACK,
		false,
	}
	types[1566] = type_
	type_ = &Type{
		"int*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1567] = type_
	type_ = &Type{
		"long",
		nil,
		T_LONG,
		KIND_STACK,
		false,
	}
	types[1568] = type_
	type_ = &Type{
		"long double",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[1569] = type_
	type_ = &Type{
		"long long",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1570] = type_
	type_ = &Type{
		"qreal",
		nil,
		T_DOUBLE,
		KIND_STACK,
		false,
	}
	types[1571] = type_
	type_ = &Type{
		"short",
		nil,
		T_SHORT,
		KIND_STACK,
		false,
	}
	types[1572] = type_
	type_ = &Type{
		"signed char",
		nil,
		T_CHAR,
		KIND_STACK,
		false,
	}
	types[1573] = type_
	type_ = &Type{
		"size_t",
		nil,
		T_ULONG,
		KIND_STACK,
		false,
	}
	types[1574] = type_
	type_ = &Type{
		"unsigned char",
		nil,
		T_UCHAR,
		KIND_STACK,
		false,
	}
	types[1575] = type_
	type_ = &Type{
		"unsigned char*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1576] = type_
	type_ = &Type{
		"unsigned int",
		nil,
		T_UINT,
		KIND_STACK,
		false,
	}
	types[1577] = type_
	type_ = &Type{
		"unsigned int*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1578] = type_
	type_ = &Type{
		"unsigned long",
		nil,
		T_ULONG,
		KIND_STACK,
		false,
	}
	types[1579] = type_
	type_ = &Type{
		"unsigned long long",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1580] = type_
	type_ = &Type{
		"unsigned short",
		nil,
		T_USHORT,
		KIND_STACK,
		false,
	}
	types[1581] = type_
	type_ = &Type{
		"va_list",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1582] = type_
	type_ = &Type{
		"void(*)()",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1583] = type_
	type_ = &Type{
		"void(*)(QObject*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1584] = type_
	type_ = &Type{
		"void(*)(QObject*,int,QAccessible::Event)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1585] = type_
	type_ = &Type{
		"void(*)(QPictureIO*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1586] = type_
	type_ = &Type{
		"void(*)(QtMsgType,const char*)",
		nil,
		T_VOIDP,
		KIND_STACK,
		false,
	}
	types[1587] = type_
	type_ = &Type{
		"void*",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1588] = type_
	type_ = &Type{
		"void**",
		nil,
		T_VOIDP,
		KIND_POINTER,
		false,
	}
	types[1589] = type_
}
