/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

/*
 * 0 ~ 299: org.gnome-settings-daemon.plugins.key-bindings
 * 300 ~ 599: org.gnome-settings-daemon.plugins.media-keys
 * 600 ~ 799: org.gnome.desktop.wm.keybindings
 * 800 ~ 899: org.compiz.shift; path: /org/compiz/profiles/shift/
 * 900 ~ 999: org.compiz.put; path: /org/compiz/profiles/put/
 */
var currentSystemBindings = map[int32]string{
	0:   "key1",
	1:   "key2",
	2:   "key3",
	3:   "key4",
	4:   "key5",
	5:   "key6",
	6:   "key7",
	7:   "key8",
	8:   "key9",
	9:   "key10",
	10:  "key11",
	645: "show-desktop",
	670: "switch-windows",
	671: "switch-windows-backward",
	302: "calculator",
	305: "eject",
	306: "email",
	334: "www",
	314: "media",
	319: "play",
	318: "pause",
	326: "stop",
	329: "volume-down",
	330: "volume-mute",
	331: "volume-up",
	320: "previous",
	315: "next",
	600: "activate-window-menu",
	601: "begin-move",
	602: "begin-resize",
	603: "close",
	611: "maximize",
	614: "minimize",
	676: "toggle-shaded",
	677: "unmaximize",
	654: "switch-to-workspace-1",
	655: "switch-to-workspace-2",
	656: "switch-to-workspace-3",
	657: "switch-to-workspace-4",
	666: "switch-to-workspace-down",
	667: "switch-to-workspace-left",
	668: "switch-to-workspace-right",
	669: "switch-to-workspace-up",
	636: "move-to-workspace-down",
	637: "move-to-workspace-left",
	638: "move-to-workspace-right",
	639: "move-to-workspace-up",
	800: "prev-key",           //switch apps with 3D
	801: "next-key",           //reverse switch apps with 3D
	900: "put-viewport-1-key", //Move window to workspace 1
	901: "put-viewport-2-key",
	902: "put-viewport-3-key",
	903: "put-viewport-4-key",
}

/* 0 ~ 299 */
var presetBindings = map[int32]string{
	0:  "key1",
	1:  "key2",
	2:  "key3",
	3:  "key4",
	4:  "key5",
	5:  "key6",
	6:  "key7",
	7:  "key8",
	8:  "key9",
	9:  "key10",
	10: "key11",
	11: "key12",
	12: "key13",
	13: "key14",
	14: "key15",
	15: "key16",
	16: "key17",
	17: "key18",
	18: "key19",
	19: "key20",
}

/* 300 ~ 599 */
var mediaBindings = map[int32]string{
	300: "area-screenshot",
	301: "area-screenshot-clip",
	302: "calculator",
	303: "capslock",
	304: "decrease-text-size",
	305: "eject",
	306: "email",
	307: "help",
	308: "home",
	309: "increase-text-size",
	310: "logout",
	311: "magnifier",
	312: "magnifier-zoom-in",
	313: "magnifier-zoom-out",
	314: "media",
	315: "next",
	316: "numlock",
	317: "on-screen-keyboard",
	318: "pause",
	319: "play",
	320: "previous",
	321: "screenreader",
	322: "screensaver",
	323: "screenshot",
	324: "screenshot-delay",
	325: "search",
	326: "stop",
	327: "terminal",
	328: "toggle-contrast",
	329: "volume-down",
	330: "volume-mute",
	331: "volume-up",
	332: "window-screenshot",
	333: "window-screenshot-clip",
	334: "www",
}

/* 600 ~ 999 */
var wmBindings = map[int32]string{
	600: "activate-window-menu",
	601: "begin-move",
	602: "begin-resize",
	603: "close",
	604: "cycle-group",
	605: "cycle-group-backward",
	606: "cycle-panels",
	607: "cycle-panels-backward",
	608: "cycle-windows",
	609: "cycle-windows-backward",
	610: "lower",
	611: "maximize",
	612: "maximize-horizontally",
	613: "maximize-vertically",
	614: "minimize",
	615: "move-to-center",
	616: "move-to-center-ne",
	617: "move-to-center-nw",
	618: "move-to-center-se",
	619: "move-to-center-sw",
	620: "move-to-side-e",
	621: "move-to-side-n",
	622: "move-to-side-s",
	623: "move-to-side-w",
	624: "move-to-workspace-1",
	625: "move-to-workspace-2",
	626: "move-to-workspace-3",
	627: "move-to-workspace-4",
	628: "move-to-workspace-5",
	629: "move-to-workspace-6",
	630: "move-to-workspace-7",
	631: "move-to-workspace-8",
	632: "move-to-workspace-9",
	633: "move-to-workspace-10",
	634: "move-to-workspace-11",
	635: "move-to-workspace-12",
	636: "move-to-workspace-down",
	637: "move-to-workspace-left",
	638: "move-to-workspace-right",
	639: "move-to-workspace-up",
	640: "panel-main-menu",
	641: "panel-run-dialog",
	642: "raise",
	643: "raise-or-lower",
	644: "set-spew-mark",
	645: "show-desktop",
	646: "switch-applications",
	647: "switch-applications-backward",
	648: "switch-group",
	649: "switch-group-backward",
	650: "switch-input-source",
	651: "switch-input-source-backward",
	652: "switch-panels",
	653: "switch-panels-backward",
	654: "switch-to-workspace-1",
	655: "switch-to-workspace-2",
	656: "switch-to-workspace-3",
	657: "switch-to-workspace-4",
	658: "switch-to-workspace-5",
	659: "switch-to-workspace-6",
	660: "switch-to-workspace-7",
	661: "switch-to-workspace-8",
	662: "switch-to-workspace-9",
	663: "switch-to-workspace-10",
	664: "switch-to-workspace-11",
	665: "switch-to-workspace-12",
	666: "switch-to-workspace-down",
	667: "switch-to-workspace-left",
	668: "switch-to-workspace-right",
	669: "switch-to-workspace-up",
	670: "switch-windows",
	671: "switch-windows-backward",
	672: "toggle-above",
	673: "toggle-fullscreen",
	674: "toggle-maximized",
	675: "toggle-on-all-workspaces",
	676: "toggle-shaded",
	677: "unmaximize",
}
